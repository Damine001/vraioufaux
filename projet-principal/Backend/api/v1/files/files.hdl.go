package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"vertoufaux/api/app/user"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Database struct {
	DB       *gorm.DB
	Enforcer *casbin.Enforcer
}

func BytesToMB(bytes uint64) float64 {
	return float64(bytes) / 1000000.0
}
func (db Database) PostFiles(ctx *gin.Context) {

	// Get the uploaded files from the request context
	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Get the user ID from the request context
	userID := ctx.Param("id")
	num, _ := strconv.Atoi(userID)
	file_theme := ctx.Param("file_theme")

	//check user id exist :
	if exists := user.CheckUserExists(db.DB, uint(num)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	// Create a new folder with the user ID
	folderPath := fmt.Sprintf("uploads/user/post/%s", userID)
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Loop through each uploaded file
	files := form.File["uploadfile"]
	for _, file := range files {

		if BytesToMB(uint64(file.Size)) > 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "File size should be less then 10MB"})
			return

		}
		// Get the file extension
		ext := filepath.Ext(file.Filename)
		allowedExts := []string{".pdf", ".doc", ".docx", ".jpg", ".jpeg", ".png", ".gif"}

		// Check if the file extension is allowed
		allowed := false
		for _, allowedExt := range allowedExts {
			if ext == allowedExt {
				allowed = true
				break
			}
		}
		if !allowed {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "File type not allowed"})
			return
		}

		// Generate a unique filename

		filename := uuid.New().String() + "_" + file.Filename

		// Open the file for reading
		src, err := file.Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		defer src.Close()

		// Create the destination file on the server
		dest, err := os.Create(fmt.Sprintf("%s/%s", folderPath, filename))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
			return
		}
		defer dest.Close()

		// Copy the uploaded file to the destination file
		_, err = io.Copy(dest, src)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// Save the file information to the database using GORM
		filepath := "uploads/user/post/" + userID + "/" + file.Filename
		fileInfo := File{
			FilePath:  filepath,
			Filename:  filename,
			FileTheme: file_theme,
			Size:      file.Size,
			UserID:    uint(num),
		}
		db.DB.Create(&fileInfo)
	}

	user_update, _ := user.GetUserByID(db.DB, uint(num))
	user_update.Points = user_update.Points + 20

	_ = user.UpdateUser(db.DB, user_update)
	// Continue with the rest of your code
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

	ctx.JSON(http.StatusOK, gin.H{"message": "Uploaded successfully"})
}

func (db Database) DownloadFile(ctx *gin.Context) {
	// Get the file name from the request parameters
	fileName := ctx.Param("fileName")

	userID := ctx.Param("id")

	num, _ := strconv.Atoi(userID)

	if exists := user.CheckUserExists(db.DB, uint(num)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid user id"})
		return
	}

	// Open the file
	filePath := "uploads/user/post/" + userID + "/" + fileName

	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	defer file.Close()

	// Get the file extension
	//ext := filepath.Ext(fileName)

	// Set the content type and attachment header
	ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+fileName)

	// Serve the file
	ctx.File(filePath)
	// Return success message
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully downloaded file"})
}

func (db Database) GetMostRecentFile(ctx *gin.Context) {

	// get id value from path

	file, err := GetMostRecentFile(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, file)
}

func (db Database) GetFiles(ctx *gin.Context) {

	// get id value from path

	file, err := GetFiles(db.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, file)
}

/*
func (db Database) DeleteFile(ctx *gin.Context) {

	// get id value from path
	file_id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Get the file name from the request parameters

	file, _ := GetFileById(db.DB, uint(file_id))

	if exists := CheckFileExists(db.DB, int64(file.ProjectID)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid file id"})
		return
	}

	// Construct the file path
	filePath := "uploads/project/" + strconv.Itoa(int(file.ProjectID)) + "/" + file.Filename // à fixer

	// Attempt to delete the file
	err_file := os.Remove(filePath)
	if err_file != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// delete form database
	if err := DeleteFile(db.DB, uint(file_id)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Return success message
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted file"})
}

func (db Database) GetFilesInProject(ctx *gin.Context) {
	projectID, err := strconv.Atoi(ctx.Param("id"))

	if exists := project.CheckProjectExists(db.DB, int64(projectID)); !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid project id"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid project ID"})
		return
	}

	files, err := GetFilesInProject(db.DB, uint(projectID))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, files)
}

*/
