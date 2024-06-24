import Link from "next/link";
import NewSecret from "./buttons/new-secret-button";

export default function Header() {
  return (
    <div className="bg-[#151515] md:w-[400px] h-[70px] p-2 flex space-x-2 rounded-lg">
      <div className=" flex flex-col justify-center w-full p-4">
        <Link href={"/"} className="font-bold text-green-300">
          Go Encrypt
        </Link>
        <p className="text-sm text-[#7f7f7f]">always keep your keys safe...</p>
      </div>
      <NewSecret />
    </div>
  );
}
