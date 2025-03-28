programa {
  inclua biblioteca Matematica --> mat
  funcao inicio() {
    real s,kW,x,z,d
    escreva("Digite o valor do salário: ")
		leia(s)
    escreva("Digite o número de Kw gasto: ")
		leia(x)
    kW= s * 0.7/100
    z= x * kW
    d= 0.9 * z
    escreva ("CUSTO POR kW:",mat.arredondar(kW,2),"\n" )
    escreva ( "CUSTO POR CONSUMO:",mat.arredondar(z,2), "\n")     
   escreva ("CUSTO COM O DESCONTO:",mat.arredondar(d,2))

  }
}
