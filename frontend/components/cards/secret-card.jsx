"use client";

import { getSecret } from "@/lib/encrypt/encrypt.api.actions";
import { useState } from "react";

export default function SecretCard() {
  const [key, setKey] = useState("");

  const handleClick = async () => {
    try {
      const data = await getSecret({ key });
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  const handleKeyChange = (e) => {
    setKey(e.target.value);
  };

  return (
    <div className="bg-[#151515] md:w-[400px] h-[400px] flex flex-col rounded-lg">
      <div className="flex p-2 rounded-lg space-x-2">
        <input
          value={key}
          onChange={handleKeyChange}
          type="text"
          placeholder="enter your key..."
          className="bg-[#242424] text-[#7f7f7f] w-full rounded-lg md:p-4 text-sm focus:"
        />

        <button
          onClick={handleClick}
          className="bg-[#242424] md:w-[70px] flex justify-center items-center rounded-lg"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="20"
            height="20"
            viewBox="0 0 24 24"
            fill="none"
            stroke="#7f7f7f"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
            class="lucide lucide-key-round"
          >
            <path d="M2 18v3c0 .6.4 1 1 1h4v-3h3v-3h2l1.4-1.4a6.5 6.5 0 1 0-4-4Z" />
            <circle cx="16.5" cy="7.5" r=".5" fill="currentColor" />
          </svg>
        </button>

        {/* <Sendkey /> */}
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
