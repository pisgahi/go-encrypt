import Link from "next/link";

export default function NewSecret() {
  return (
    <Link
      href={"/add"}
      className="bg-green-300 w-[35px] md:w-[70px] flex justify-center items-center rounded-lg"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        width="20"
        height="20"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
        class="lucide lucide-lock"
      >
        <rect width="18" height="11" x="3" y="11" rx="2" ry="2" />
        <path d="M7 11V7a5 5 0 0 1 10 0v4" />
      </svg>
    </Link>
  );
}
