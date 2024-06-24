export default function Sendkey() {
  return (
    <button className="bg-[#242424] md:w-[70px] flex justify-center items-center rounded-lg">
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
      </svg>{" "}
    </button>
  );
}
