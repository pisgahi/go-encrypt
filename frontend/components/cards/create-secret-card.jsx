"use client";

import { createSecret } from "@/lib/encrypt/encrypt.api.actions";
import { useState } from "react";

export default function CreateSecret() {
  const [key, setKey] = useState("");
  const [plainText, setPlainText] = useState("");

  const handleClick = async () => {
    try {
      const data = await createSecret({ plainText, key });
    } catch (error) {
      console.error("Error creating secret:", error);
    }
  };

  const handleKeyChange = (e) => {
    setKey(e.target.value);
  };

  const handlePlainTextChange = (e) => {
    setPlainText(e.target.value);
  };

  return (
    <div>
      <div className="bg-[#151515] w-[350px] md:w-[400px] p-2 rounded-lg">
        <div className="flex pb-2 rounded-lg">
          <input
            value={key}
            onChange={handleKeyChange}
            type="text"
            placeholder="your 32-bit key"
            className="bg-[#242424] text-[#7f7f7f] w-full rounded-lg p-4 text-xs focus:outline-none"
          />
        </div>
        <div className="bg-[#242424] h-[320px] rounded-lg">
          <textarea
            value={plainText}
            onChange={handlePlainTextChange}
            placeholder="your secret goes here..."
            className="bg-[#242424] text-[#7f7f7f] text-sm resize-none 
          w-full h-full rounded-lg p-4 mb-6 focus:outline-none"
          />
        </div>
      </div>
      <div
        className="bg-[#242424] hover:bg-green-300 transition duration-350 
        ease-in-out flex items-center p-2 rounded-lg mt-12 group"
      >
        <button
          onClick={handleClick}
          className="w-full p-1 flex justify-center items-center rounded-lg"
        >
          <p className=" text-[#7f7f7f] group-hover:text-white">
            Create Secret
          </p>
        </button>
      </div>
    </div>
  );
}
