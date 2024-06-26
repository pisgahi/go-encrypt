import GithubButton from "@/components/buttons/github-button";
import SecretCard from "@/components/cards/secret-card";

export default function Home() {
  return (
    <div>
      <SecretCard />
      <div className="mt-10">
        <GithubButton />
      </div>
    </div>
  );
}
