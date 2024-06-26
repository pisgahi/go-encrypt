import GithubButton from "./buttons/github-button";
import SecretCard from "./cards/secret-card";

export default function Homepage() {
  return (
    <div>
      <SecretCard />
      <div className="mt-10">
        <GithubButton />
      </div>
    </div>
  );
}
