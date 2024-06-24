export default function CreateSecret() {
  return (
    <div className="bg-[#151515] md:w-[400px] p-2 rounded-lg">
      <div className="flex pb-2 rounded-lg">
        <input
          type="text"
          placeholder="your 32-bit key"
          className="bg-[#242424] text-[#7f7f7f] w-full rounded-lg md:p-4 text-xs focus:outline-none"
        />
      </div>

      <div className="bg-[#242424] h-[320px] rounded-lg">
        <textarea
          placeholder="your secret goes here..."
          className="bg-[#242424] text-[#7f7f7f] resize-none 
          w-full h-full rounded-lg p-4 mb-6 focus:outline-none"
        />
      </div>

      <div
        className="bg-[#242424] hover:bg-green-300 transition duration-350 
        ease-in-out flex items-center p-2 rounded-lg mt-2 group"
      >
        <button className="w-full p-1 flex justify-center items-center rounded-lg">
          <p className=" text-[#7f7f7f] group-hover:text-white">
            Create Secret
          </p>
        </button>
      </div>
    </div>
  );
}
