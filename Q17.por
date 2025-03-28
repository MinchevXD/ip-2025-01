programa {
  funcao inicio() {
     real x,y,i,z
   escreva ("DIGITE DOIS VALORES:")
   leia (x,y) 
   se (x%2==0)
   para (i=0;i<=y-1;i++)
   escreva (x+i*2," ")
   senao 
   escreva ("O PRIMEIRO NUMERO NAO E PAR","\n")
  }
}
