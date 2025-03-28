programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real a,b,x
    escreva ("DIGITE O SALÁRIO ATUAL:")
    leia (a)
    se (a<=300)
    escreva ("SALARIO COM REAJUSTE=",mat.arredondar(a+0.5*a,2))
    senao 
    escreva ("SALARIO COM REAJUSTE=",mat.arredondar(a+0.3*a,2))
  }
}
