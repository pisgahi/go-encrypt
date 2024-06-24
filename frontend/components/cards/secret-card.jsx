import Sendkey from "../buttons/send-key-button";

export default function SecretCard() {
  return (
    <div className="bg-[#151515] md:w-[400px] h-[400px] flex flex-col rounded-lg">
      <div className="flex p-2 rounded-lg space-x-2">
        <input
          type="text"
          placeholder="enter your key..."
          className="bg-[#242424] text-[#7f7f7f] w-full rounded-lg md:p-4 text-sm focus:"
        />
        <Sendkey />
      </div>
      <div className=" w-full h-full rounded-lg px-2 pb-2">
        <div
          className="flex flex-col items-center 
          relative overflow-hidden"
        >
          <div className="bg-[#242424] h-[320px] overflow-y-auto rounded-lg">
            <p className=" text-[#7f7f7f] h-full rounded-lg p-4 mb-6">
              secret info Lorem ipsum dolor sit amet consectetur adipisicing
              elit. Itaque, dignissimos esse consectetur nemo qui eius doloribus
              explicabo, facilis natus voluptates aperiam expedita fuga
              aspernatur fugiat similique neque assumenda, adipisci odit? Lorem
              ipsum, dolor sit amet consectetur adipisicing elit. Totam sint
              nobis magnam deserunt perferendis esse doloremque amet dolor
              aperiam illum! Magnam consequuntur odio non officia voluptatibus
              voluptatem ipsum magni. Nihil! Lorem ipsum dolor sit amet
            </p>
          </div>
        </div>
      </div>
    </div>
  );
}
