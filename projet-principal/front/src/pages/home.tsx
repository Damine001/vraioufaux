import Navbar from "@/components/navbar";
import RankingCharts from "@/components/ranking";
import CountUp from "react-countup";

export default function Home() {
  return (
    <div>
      <div className="h-[92vh] bg-black grid place-content-center">
        <div className="flex items-center">
          <img src="./src/assets/logo.svg" alt="logo" className="w-36 h-36" />
          <h1 className="text-white text-7xl">Vert ou faux</h1>
        </div>
      </div>
      <Navbar className="h-14 sticky top-0" />
      <div className="fixed bottom-9 left-9 text-red-600 animate-bounce z-50">
        Cherche Charlie !
      </div>
      <div className="container grid grid-cols-12 mt-16 gap-x-8 gap-y-12 px-16 items-center">
        <div className="col-span-7">
          <h2 className="text-4xl font-bold mb-4">
            Changement Climatique: Faire le Tri entre Fausse et Vraie
            Information
          </h2>
          <br />
          <p>
            Dans un monde où l'information est abondante, il est crucial de
            distinguer la vérité de la désinformation concernant le changement
            climatique. Découvrez comment reconnaître les sources fiables et les
            faits scientifiques. Rejoignez-nous dans notre combat contre la
            désinformation climatique.
          </p>
          <br />
          <button className="bg-red-600 text-white px-4 py-2 rounded-md">
            Agisser maintenant
          </button>
        </div>
        <div className="col-span-5">
          <img
            src="./src/assets/charlie3.png"
            alt="fake charlie"
            className="w-24 relative bottom-[-5rem]"
          />
          <img
            src="./src/assets/image1.webp"
            alt=""
            className="h-auto max-w-full rounded-3xl relative z-10"
          />
        </div>
        <div className="col-span-5">
          <img
            src="./src/assets/charlie2.png"
            alt="fake charlie"
            className="w-24 relative bottom-[-9rem] right-[-27.5rem] rotate-90"
          />
          <img
            src="./src/assets/image2.jpg"
            alt=""
            className="h-auto max-w-full rounded-3xl relative z-10"
          />
        </div>
        <div className="col-span-7">
          <h2 className="text-4xl font-bold mb-4">
            Les Conséquences du Changement Climatique
          </h2>
          <br />
          <div className="grid grid-cols-3 gap-4">
            <div>
              <h3 className="text-xl mb-4">Réchauffement de la Planète</h3>
              <p>
                Les températures mondiales augmentent, ce qui entraîne des
                vagues de chaleur, des tempêtes plus fortes et des conditions
                climatiques extrêmes.
              </p>
            </div>
            <div>
              <h3 className="text-xl mb-4">Fonte des Glaces</h3>
              <p>
                Les calottes glaciaires et les glaciers fondent, ce qui entraîne
                une élévation du niveau de la mer et des inondations côtières.
              </p>
            </div>
            <div>
              <h3 className="text-xl mb-4">Perturbation de la Biodiversité</h3>
              <p>
                Le changement climatique affecte les habitats naturels, mettant
                ainsi en péril de nombreuses espèces végétales et animales. Les
                Conséquences du Changement Climatique
              </p>
            </div>
          </div>
        </div>
      </div>
      <div className="container my-16 px-16">
        <h2 className="text-3xl mb-16">
          L'Importance du Tri pour l'Action Climatique
        </h2>
        <p>
          Comprenez l'impact de la désinformation sur les mesures prises pour
          lutter contre le changement climatique. Explorez le rôle essentiel du
          public dans la diffusion de l'information vérifiée et découvrez
          comment votre participation peut aider à promouvoir des actions
          climatiques concrètes et efficaces.
        </p>
      </div>
      <div className="h-40 my-16 mx-16 flex justify-around items-center">
        <div className="border p-8 grid place-items-center bg-gray-100 shadow-md rounded-2xl text-2xl">
          <span>
            +<CountUp end={20} /> comptes
          </span>
        </div>
        <div className="border p-8 grid place-items-center bg-gray-100 shadow-md rounded-2xl text-2xl">
          <span>
            +<CountUp end={30} /> certifcations accordées
          </span>
        </div>
        <div className="border p-8 grid place-items-center bg-gray-100 shadow-md rounded-2xl text-2xl">
          <span>
            +<CountUp end={5} /> idées innovantes
          </span>
        </div>
      </div>
      <div className="h-[20rem] w-[30rem] mx-auto">
        <img
          src="./src/assets/charlie.png"
          alt="fake charlie"
          className="w-16 relative bottom-[-16rem] right-[-4rem] rotate-12"
        />
        <RankingCharts />
      </div>
    </div>
  );
}
