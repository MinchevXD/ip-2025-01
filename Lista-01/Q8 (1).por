programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real r, h,a,p
 
    escreva ("escreva o valor do raio:")
    leia (r)
    escreva ("escreva o valor da altura:")
    leia (h)
    a= (3.14159 * r * r * 2)+ (2 * 3.14159 * r * h ) 
    p= a * 100
    escreva ("O VALOR DO CUSTO E=",mat.arredondar (p,2),"\n")
  }
}