import Image from "next/image";
import Link from "next/link";

export function Navbar() {
  return (
    <div className="flex max-w-full items-center justify-items-stretch rounded-2xl bg-[#1D232A] px-6 py-2 shadow-nav">
      <div className="flex grow items-center justify-center">
        <Link href="/"> 
          <Image
<<<<<<< HEAD
            src="/icone.svg"
=======
            src="/nextjs/icone.svg"
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
            alt="Icone DevTicket"
            width={136}
            height={48}
            className="max-h-[48px]"
          />
        </Link>
      </div>
      <Link href={"/checkout"} className="min-h-6 min-w-6 grow-0 items-center">
        <Image
<<<<<<< HEAD
          src="/cart-outline.svg"
=======
          src="/nextjs/cart-outline.svg"
>>>>>>> 3febb45 (Mudanças nos docker-compose das pastas referentes ao nest, next e golang para rodar o projeto como um todo. Criação do readme.md final para melhor compreensão da execução do projeto.)
          alt="Icone de carrinho"
          width={24}
          height={24}
        />
      </Link>
    </div>
  );
}
