programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real a,h,v
    escreva ("DIGITE O VALOR DA ALTURA E ARESTA DA PIRÂMIDE RESPECTIVAMENTE:","\n")
    leia (a , h)
    v=h*a*a*mat.raiz (3, 2.0)/2
    escreva ("O VOLUME DA PIRAMIDE E =",mat.arredondar(v,2),"  METROS CUBICOS")
  }
}
