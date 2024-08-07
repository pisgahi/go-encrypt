import { Inter } from "next/font/google";
import "./globals.css";
import Header from "@/components/header";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Go Encrypt",
  description: "Encrypt your secrets",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <div className="flex flex-col items-center pt-10 space-y-2">
          <Header />
          {children}
        </div>
      </body>
    </html>
  );
}
