package controller

import (
	"net/http" // Uncomment this line
	"github.com/gin-gonic/gin"
	"example.com/gogin/config"
	"example.com/gogin/models"


"gorm.io/gorm"

"sync"





)


func Calculo(c *gin.Context) {
	// Cria um canal para sincronização
	done := make(chan bool)

	// Utiliza WaitGroup para esperar a conclusão de todas as goroutines
	var wg sync.WaitGroup

	// Número de goroutines que serão iniciadas
	numGoroutines := 16

	// Adiciona o número de goroutines ao WaitGroup
	wg.Add(numGoroutines)

	// Inicia as goroutines
	go worker(&wg, done,
		SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
	
		// Aguarda a conclusão de todas as goroutines
		go worker(&wg, done,
			SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
			go worker(&wg, done,
				SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
				go worker(&wg, done,
					SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
					go worker(&wg, done,
						SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
						go worker(&wg, done,
							SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
							go worker(&wg, done,
								SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
								go worker(&wg, done,
									SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
									go worker(&wg, done,
										SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
										go worker(&wg, done,
											SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
											go worker(&wg, done,
												SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
												go worker(&wg, done,
													SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
													go worker(&wg, done,
														SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
														go worker(&wg, done,
															SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
															go worker(&wg, done,
																SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																go worker(&wg, done,
																	SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																	go worker(&wg, done,
																		SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)	


																		go worker(&wg, done,
																			SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																		
																			// Aguarda a conclusão de todas as goroutines
																			go worker(&wg, done,
																				SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
																				go worker(&wg, done,
																					SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
																					go worker(&wg, done,
																						SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																						go worker(&wg, done,
																							SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																							go worker(&wg, done,
																								SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																								go worker(&wg, done,
																									SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																									go worker(&wg, done,
																										SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																										go worker(&wg, done,
																											SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																											go worker(&wg, done,
																												SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																												go worker(&wg, done,
																													SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																													go worker(&wg, done,
																														SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																														go worker(&wg, done,
																															SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																															go worker(&wg, done,
																																SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																go worker(&wg, done,
																																	SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																	go worker(&wg, done,
																																		SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																		go worker(&wg, done,
																																			SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)				
																																			
																																			go worker(&wg, done,
																																				SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																			
																																				// Aguarda a conclusão de todas as goroutines
																																				go worker(&wg, done,
																																					SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
																																					go worker(&wg, done,
																																						SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
																																						go worker(&wg, done,
																																							SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																							go worker(&wg, done,
																																								SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																								go worker(&wg, done,
																																									SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																									go worker(&wg, done,
																																										SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																										go worker(&wg, done,
																																											SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																											go worker(&wg, done,
																																												SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																												go worker(&wg, done,
																																													SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																													go worker(&wg, done,
																																														SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																														go worker(&wg, done,
																																															SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																															go worker(&wg, done,
																																																SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																go worker(&wg, done,
																																																	SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																	go worker(&wg, done,
																																																		SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																		go worker(&wg, done,
																																																			SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																			go worker(&wg, done,
																																																				SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)	
																																		
																																		
																																																				go worker(&wg, done,
																																																					SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																				
																																																					// Aguarda a conclusão de todas as goroutines
																																																					go worker(&wg, done,
																																																						SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
																																																						go worker(&wg, done,
																																																							SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
																																																							go worker(&wg, done,
																																																								SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																								go worker(&wg, done,
																																																									SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																									go worker(&wg, done,
																																																										SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																										go worker(&wg, done,
																																																											SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																											go worker(&wg, done,
																																																												SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																												go worker(&wg, done,
																																																													SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																													go worker(&wg, done,
																																																														SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																														go worker(&wg, done,
																																																															SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																															go worker(&wg, done,
																																																																SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																																go worker(&wg, done,
																																																																	SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																																	go worker(&wg, done,
																																																																		SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																																		go worker(&wg, done,
																																																																			SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																																			go worker(&wg, done,
																																																																				SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)
																																																																				go worker(&wg, done,
																																																																					SomaAdicionalno, SomaAuxilioAlimentacao, SomaCalculoHoraExtra, SomaGratificacao, SomaInss, SomaPericulosidade, Somarh, SomaSalarioBase)		
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()
	go func() {
		wg.Wait()
		close(done)
	}()


	// Aguarda o canal done ser fechado
	<-done

	// Agora, você pode continuar com o resto do código, como enviar a resposta JSON
	c.JSON(http.StatusOK, gin.H{"resultado": "o pai TÁ on"})
	break
}

func worker(wg *sync.WaitGroup, done chan bool, calculos ...func()) {
	defer wg.Done()

	// Chama todas as funções de cálculo
	for _, calculo := range calculos {
		calculo()
		
	}

	done <- true
}


func Somarh(){

	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)


	somaPorRhUnidadee := make(map[int]map[string]int)
	
	
	for _, pdt := range pdts {
		if _, ok := somaPorRhUnidadee[pdt.Rh]; !ok {
			somaPorRhUnidadee[pdt.Rh] = make(map[string]int)
		}
		somaPorRhUnidadee[pdt.Rh][pdt.Unidade] += 1
	}


	for _, pdt := range pdts {
		soma := somaPorRhUnidadee[pdt.Rh][pdt.Unidade]
		config.DB.Model(&pdt).Update("Rh", pdt.Rh).Update("Unidade", pdt.Unidade).Update("Rh", soma)
	}


	
	somaPorRhUnidadeee := make(map[int]map[string]int)

	for _, iadvh := range iadvhs {
		if _, ok := somaPorRhUnidadeee[iadvh.Rh]; !ok {
			somaPorRhUnidadeee[iadvh.Rh] = make(map[string]int)
		}
		somaPorRhUnidadeee[iadvh.Rh][iadvh.Unidade] += 1
	}

	
	for _, iadvh := range iadvhs {
		soma := somaPorRhUnidadeee[iadvh.Rh][iadvh.Unidade]
		config.DB.Model(&iadvh).Update("Rh", iadvh.Rh).Update("Unidade", iadvh.Unidade).Update("Rh", soma)
	}



	var iadvhResults []models.Iadvh
	config.DB.Find(&iadvhResults)

	var pdtResults []models.PDT
	config.DB.Find(&pdtResults)


	var results []struct {
		Rh      int
		Unidade string
		Funcao  string
		//Adicional_noturnoo int
	}
	for _, iadvh := range iadvhResults {
		results = append(results, struct {
			Rh      int
			Unidade string
			Funcao  string
			//Adicional_noturnoo int
		}{
			Rh:      iadvh.Rh,
			Unidade: iadvh.Unidade,
			Funcao:  iadvh.Funcao,
			//Adicional_noturnoo: iadvh.Adicional_noturnoo,
		})
	}
	for _, pdt := range pdtResults {
		results = append(results, struct {
			Rh      int
			Unidade string
			Funcao  string
			//Adicional_noturnoo int

		}{
			Rh:      pdt.Rh,
			Unidade: pdt.Unidade,
			Funcao:  pdt.Funcao,
			//Adicional_noturnoo: pdt.Adicional_noturnoo,
		})
	}

	
	somaPorRhUnidade := make(map[int]map[string]int)

	for _, result := range results {
		if _, ok := somaPorRhUnidade[result.Rh]; !ok {
			somaPorRhUnidade[result.Rh] = make(map[string]int)
		}
		somaPorRhUnidade[result.Rh][result.Unidade] += 1
	}

	for _, result := range results {
		soma := somaPorRhUnidade[result.Rh][result.Unidade]
		config.DB.Model(&models.PDT{}).
			Where("Unidade = ? AND Funcao = ?", result.Unidade, result.Funcao).
			UpdateColumn("Rh", gorm.Expr("Rh + ?", soma))

	         i:= 1
			if i==1{
				break
			}
}



}


/////////////////////////////////////////////////////////
//função de adicional noturno

func SomaAdicionalno(){

	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)

somaPoradicionalnoturnoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20]; !ok {
			somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20] = make(map[string]map[string]float64)
		}
		if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade]; !ok {
			somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade] = make(map[string]float64)
		}
		somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao] += pdt.Adicional_noturno_20+pdt.Adicional_noturno_30+pdt.Adicional_noturno_20_valor+pdt.Media_adicional_noturno_salario_maternidade+pdt.Adicional_noturno_30_valor+pdt.Adicional_noturno_30_mes_anterior+pdt.Diferenca_adicional_noturno+pdt.Adicional_noturno_20_retroativo
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior]; !ok {
			somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior] = make(map[string]map[string]float64)
		}
		if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade]; !ok {
			somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade] = make(map[string]float64)
		}
		somaPoradicionalnoturnoUnidadeFuncaoPdt[iadvh.Adicional_noturno_mes_anterior][iadvh.Unidade][iadvh.Funcao] += iadvh.Adicional_noturno+iadvh.Adicional_noturno_mes_anterior
	}

	

	for _, pdt := range pdts {
		soma := somaPoradicionalnoturnoUnidadeFuncaoPdt[pdt.Adicional_noturno_20][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Adicional_noturnofo", soma)
	}


}








/////////////////////////////////////////////////////////////////////////
//calculoauxilio alimentação
func SomaAuxilioAlimentacao(){


	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)



	somaPorauxilioalimentacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade] = make(map[string]float64)
		}
		somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade][pdt.Funcao] += pdt.Auxilio_alimentacao+pdt.Auxilio_alimentacao_retroativo
	}

	

	for _, pdt := range pdts {
		soma := somaPorauxilioalimentacaoUnidadeFuncaoPdt[pdt.Auxilio_alimentacao_retroativo][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Auxilio_alimentacaofo", soma)
	}


}



/////////////////////////////////////////////////////////////////////////
//calculo periculosidade
func SomaPericulosidade(){


	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)







	somaPorpericulosidadeUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade] = make(map[string]float64)
		}
		somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade][pdt.Funcao] += pdt.Periculosidade
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade]; !ok {
			somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorpericulosidadeUnidadeFuncaoPdt[iadvh.Periculosidade][iadvh.Unidade][iadvh.Funcao] += iadvh.Periculosidade
	}

	

	for _, pdt := range pdts {
		soma := somaPorpericulosidadeUnidadeFuncaoPdt[pdt.Periculosidade][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Periculosidadefo", soma)
	}


}






///////////////////////////////////////////////////////////////////////////
//calculo gratificação
func SomaGratificacao(){


	
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)



	

	somaPorgratificacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade] = make(map[string]float64)
		}
		somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_funcao_retroativo+pdt.Gratificacao_funcao+pdt.Gratificacao_dedicacao_exclusiva+pdt.Gratificacao_produtividade+pdt.Gratificacao_funcao_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermage+pdt.Gratificacao_5_sobre_salario_base+pdt.Prorrogacao_media_deducao_exclusiva_salario_maternidade+pdt.Gratificacao_funcao_instrumentacao_cirurgico+pdt.Gratificacao_por_interiorizacao+pdt.Gratificacao+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_funcao_prorrogacao_maternidade+pdt.Gratificacao_funcao_supervisor_enfermagem_sl
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade]; !ok {
			somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorgratificacaoUnidadeFuncaoPdt[iadvh.Gratificacao_funcao_instrumentacao_cirurgico][iadvh.Unidade][iadvh.Funcao] += iadvh.Gratificacao+iadvh.Gratificacao_cargo_confianca+iadvh.Gratificacao_indenizatoria+iadvh.Gratificacao_funcao_instrumentacao_cirurgico+iadvh.Gratificacao_mes_anterior
	}

	

	for _, pdt := range pdts {
		soma := somaPorgratificacaoUnidadeFuncaoPdt[pdt.Gratificacao_funcao_instrumentacao_cirurgico][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Gratificacaofo", soma)
	}

}





////////////////////////////////////////////////////////////////////////////
//calculo hora extra
func SomaCalculoHoraExtra(){

	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)




	

	somaPorhoraextraUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade] = make(map[string]float64)
		}
		somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade][pdt.Funcao] += pdt.Media_horas_extras_prorrogacao_salario_maternidade+pdt.Hora_extra_50+pdt.Hora_extra_sobre_adicionais_50+pdt.Hora_extra_sobre_adicionais_60+pdt.Hora_extra_sobre_adicionais_100+pdt.Dsr_sobre_horas_extras+pdt.Dsr_sem_horas_extras_retroativas+pdt.Hora_extra_100+pdt.Hora_extra_retroativa_60+pdt.Hora_extra_50_valor+pdt.Hora_extra_100_mes_anterior+pdt.Hora_extra_50_retroativo+pdt.Hora_extra_100_valor+pdt.Hora_extra_60_valor
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade]; !ok {
			somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorhoraextraUnidadeFuncaoPdt[iadvh.Hora_extra_50][iadvh.Unidade][iadvh.Funcao] += iadvh.Hora_extra_100+iadvh.Hora_extra_50+iadvh.Hora_extra_60+iadvh.Hora_extra_60_mes_anterior+iadvh.Horas_extras_100_mes_anterior+iadvh.Horas_extras_50_mes_anterior
	}

	

	for _, pdt := range pdts {
		soma := somaPorhoraextraUnidadeFuncaoPdt[pdt.Hora_extra_50][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Hora_extrafo", soma)
	}

}



///////////////////////////////////////////////////////////////////////////
//calculo salario base
func SomaSalarioBase(){

	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)





	somaPorsalariobaseUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

	for _, pdt := range pdts {
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade] = make(map[string]float64)
		}
		somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao] += pdt.Salario_base+pdt.Jovem_aprendiz+pdt.Diferenca_salarial+pdt.Devolucao_falta+pdt.Reembolso_faltas+pdt.Salario_substituicao+pdt.Salario_substituicao_mes_anterior+pdt.Estorno_desconto_vale_transporte+pdt.Dias_atestado+pdt.Prorrogacao_licenca_maternidade+pdt.Licenca_remunerada+pdt.Reembolso_educacional+pdt.Insuficiencia_saldo_provento+pdt.Saldo_salario+pdt.Saldo_salario_horista
	}


	for _, iadvh := range iadvhs {
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade]; !ok {
			somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorsalariobaseUnidadeFuncaoPdt[iadvh.Saldo_salario][iadvh.Unidade][iadvh.Funcao] += iadvh.Atestado_medico+iadvh.Devolucao_desconto_indevido+iadvh.Devolucao_faltas+iadvh.Desconto_plano_saude_hapvida+iadvh.Desconto_plano_odontologico_integral+iadvh.Dias_trabalhados+iadvh.Diferenca_salarial+iadvh.Horas_noturnas_reduzidas+iadvh.Insuficiencia_saldo_proventos+iadvh.Licenca_paternidade+iadvh.Saldo_salario
	}

	
	for _, pdt := range pdts {
		soma := somaPorsalariobaseUnidadeFuncaoPdt[pdt.Saldo_salario][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Salariobasefo", soma)
	}

	
}



	
//////////////////////////////////////////////////////////////////////////
// calculo inss
func SomaInss(){


	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)









	somaPordeducaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)

	for _, pdt := range pdts {
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade] = make(map[string]float64)
		}
		somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado  //aqui tem 12
	}


	for _, iadvh := range iadvhs {
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade] = make(map[string]map[string]float64)
		}
		if _, ok := somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade]; !ok {
			somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade] = make(map[string]float64)
		}
		somaPordeducaoUnidadeFuncaoPdt[iadvh.Salario_maternidade][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade
	}

	
	for _, pdt := range pdts {
		soma := somaPordeducaoUnidadeFuncaoPdt[pdt.Salario_maternidade][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Deducao_inssfo", soma)
	}




}


func GetUserss(context *gin.Context) {
    var pdt []models.PDT
	var pdts []models.PDT
	config.DB.Find(&pdts)

	var iadvhs []models.Iadvh
	config.DB.Find(&iadvhs)

//////////////////////////////////////////////////////////////////////////
//calculo encargos

//falta fazer dsr 


	somaPorencargosUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


	for _, pdt := range pdts {
		if _, ok := somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado]; !ok {
			somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade]; !ok {
			somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade] = make(map[string]float64)
		}
		somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao] += (pdt.Fgts_folha+pdt.Ferias+pdt.Ferias_nao_gozadas+pdt.Indenizacao_ferias_nao_gozadas+pdt.Ferias_sobre_aviso_indenizado+pdt.Aviso_previo_indenizado+pdt.Decimo_terceiro_sobre_aviso_indenizado+pdt.Ferias_vencidas+pdt.Ferias_proporcionais+pdt.Fgts_recisao+pdt.Fgts_decimo_terceiro_salario_recisao+pdt.Um_terco_ferias_recisao+pdt.Um_terco_ferias+pdt.Fgts_ferias) - (pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado )
	}

	for _, iadvh := range iadvhs {
		if _, ok := somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado]; !ok {
			somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado] = make(map[string]map[string]float64)
		}
		if _, ok := somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade]; !ok {
			somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade] = make(map[string]float64)
		}
		somaPorencargosUnidadeFuncaoPdt[iadvh.Aviso_previo_indenizado][iadvh.Unidade][iadvh.Funcao] += (iadvh.Um_terco_ferias_mes+iadvh.Um_terco_ferias_proximo_mes+iadvh.Decimo_terceiro_indenizado+iadvh.Decimo_terceiro_licenca_maternidade+iadvh.Decimo_terceiro_recisao+iadvh.Adicional_um_terco_ferias+iadvh.Adicional_um_terco_ferias_proporcional_recisao+iadvh.Aviso_previo_indenizado+iadvh.Ferias_indenizadas+iadvh.Ferias_mes+iadvh.Ferias_proximo_mes+iadvh.Ferias_proporcionais+iadvh.Media_decimo_terceiro_rescisao+iadvh.Media_aviso_previo+iadvh.Media_ferias+iadvh.Media_ferias_proporcionais+iadvh.Multa_art_479_clt)- ( iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade)
	}

	

	for _, pdt := range pdts {
		soma := somaPorencargosUnidadeFuncaoPdt[pdt.Aviso_previo_indenizado][pdt.Unidade][pdt.Funcao]
		config.DB.Model(&pdt).Update("Encargosfo", soma)
	}

	context.JSON(200, gin.H{
		"message": "Soma Encargosfo atualizada na tabela pdt com sucesso!",
	})
//////////////////////////////////////////////////////////////////////////
//calculo insalubridade


		
			somaPorInsalubridadeUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
			for _, pdt := range pdts {
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40] = make(map[string]map[string]float64)
				}
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade] = make(map[string]float64)
				}
				somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade][pdt.Funcao] += pdt.Insalubridade_40 + pdt.Insalubridade_sem_base_40+pdt.Insalubridade_sem_minimo_20+ pdt.Insalubridade_sem_minimo_20+pdt.Insalubridade_retroativa+pdt.Media_insalubridade_prorrogacao_salario_maternidade
			}
		
			
			for _, iadvh := range iadvhs {
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40] = make(map[string]map[string]float64)
				}
				if _, ok := somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade]; !ok {
					somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade] = make(map[string]float64)
				}
				somaPorInsalubridadeUnidadeFuncaoPdt[iadvh.Insalubridade_40][iadvh.Unidade][iadvh.Funcao] += iadvh.Diferenca_insalubridade+ iadvh.Insalubridade_20+ iadvh.Insalubridade_40+iadvh.Insalubridade_mes_anterior_40
			}
		
			
	
			for _, pdt := range pdts {
				soma := somaPorInsalubridadeUnidadeFuncaoPdt[pdt.Insalubridade_40][pdt.Unidade][pdt.Funcao]
				config.DB.Model(&pdt).Update("Insalubridade", soma)
			}
		
			context.JSON(200, gin.H{
				"message": "Soma Insalubridade_40 e Gratificacaofo atualizada na tabela pdt com sucesso!",
			})

/////////////////////////////////////////////////////////////////////////////
//total salario folha
somaPortotalsalariofolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
			for _, pdt := range pdts {
				if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha]; !ok {
					somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade]; !ok {
					somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade] = make(map[string]float64)
				}
				somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao] += pdt.Salariobasefo+pdt.Gratificacaofo+pdt.Insalubridade+pdt.Periculosidadefo+pdt.Adicional_noturnofo+pdt.Dsrfo
			}
		
			
			for _, iadvh := range iadvhs {
				if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha]; !ok {
					somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade]; !ok {
					somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade] = make(map[string]float64)
				}
				somaPortotalsalariofolhaUnidadeFuncaoPdt[iadvh.Total_salario_folha][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_base+iadvh.Gratificaoia+iadvh.Insalubridade+iadvh.Periculosidadeia+iadvh.Adicional_noturnoia+iadvh.Dsria
			}
		
			
	
			for _, pdt := range pdts {
				soma := somaPortotalsalariofolhaUnidadeFuncaoPdt[pdt.Total_salario_folha][pdt.Unidade][pdt.Funcao]
				config.DB.Model(&pdt).Update("Total_salario_folha", soma)
			}
		
			context.JSON(200, gin.H{
				"message": "Soma Insalubridade_40 e Gratificacaofo atualizada na tabela pdt com sucesso!",
			})
////////////////////////////////////////////////////////////////////////////////////////////
//total salario mensal

			somaPortotalsalariomensalfolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
			for _, pdt := range pdts {
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade] = make(map[string]float64)
				}
				somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao] +=1// somaPorencargosUnidadeFuncaoPdt + somaPortotalsalariofolhaUnidadeFuncaoPdt
			}
		
			
			for _, iadvh := range iadvhs {
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha] = make(map[string]map[string]float64)
				}
				if _, ok := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade]; !ok {
					somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade] = make(map[string]float64)
				}
				somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[iadvh.Total_mensal_folha][iadvh.Unidade][iadvh.Funcao] +=1 //iadvh.Encargosia + iadvh.Total_salario_folha
			}
		
			
	
			for _, pdt := range pdts {
				soma := somaPortotalsalariomensalfolhaUnidadeFuncaoPdt[pdt.Total_mensal_folha][pdt.Unidade][pdt.Funcao]
				config.DB.Model(&pdt).Update("Total_mensal_folha", soma)
			}
		
			context.JSON(200, gin.H{
				"message": "Soma Insalubridade_40 e Gratificacaofo atualizada na tabela pdt com sucesso!",
			})
//////////////////////////////////////////////////////////////////////////////
// total salario anual

somaPortotalsalarioanualfolhaUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)
		
		
for _, pdt := range pdts {
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade] = make(map[string]float64)
	}
	somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao] += (pdt.Total_mensal_folha)*12
}


for _, iadvh := range iadvhs {
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha] = make(map[string]map[string]float64)
	}
	if _, ok := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade]; !ok {
		somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade] = make(map[string]float64)
	}
	somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[iadvh.Total_anual_folha][iadvh.Unidade][iadvh.Funcao] += (iadvh.Total_mensal_folha)*12
}



for _, pdt := range pdts {
	soma := somaPortotalsalarioanualfolhaUnidadeFuncaoPdt[pdt.Total_anual_folha][pdt.Unidade][pdt.Funcao]
	config.DB.Model(&pdt).Update("Total_anual_folha", soma)
}

context.JSON(200, gin.H{
	"message": "Soma Insalubridade_40 e Gratificacaofo atualizada na tabela pdt com sucesso!",
})



/////////////////////////////////////////////////////////////////////////////
//post



    var requestData struct {
        Unidade string `json:"unidade" binding:"required"`
    }

	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, forneça a unidade no corpo da solicitação"})
		return
	}



	if err := config.DB.Where("Unidade = ?", requestData.Unidade).Find(&pdt).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar o banco de dados"})
		return
	}

	if len(pdt) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"error": "Nenhum registro encontrado para a unidade especificada"})
		return
	}

	// Usando um mapa para armazenar funções únicas
	funcoesUnicas := make(map[string]struct{})

	// Filtra funções únicas e atualiza o slice pdt
	var pdtResultado []models.PDT
	for _, item := range pdt {
		if _, ok := funcoesUnicas[item.Funcao]; !ok {
			funcoesUnicas[item.Funcao] = struct{}{}
			pdtResultado = append(pdtResultado, item)
		}
	}

	context.JSON(http.StatusOK, pdtResultado)

   /* // Analisa o JSON do corpo da solicitação e verifica se o campo "unidade" está presente
    if err := context.ShouldBindJSON(&requestData); err != nil {
        context.JSON(http.StatusBadRequest, gin.H{"error": "Por favor, forneça a unidade no corpo da solicitação"})
        return
    }

    // Filtra os resultados com base na unidade
    if err := config.DB.Where("Unidade = ?", requestData.Unidade).Find(&pdt).Error; err != nil {
        context.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar o banco de dados"})
        return
    }

    // Verifica se foram encontrados registros
    if len(pdt) == 0 {
        context.JSON(http.StatusNotFound, gin.H{"error": "Nenhum registro encontrado para a unidade especificada"})
        return
    }

    context.JSON(http.StatusOK, pdt)*/
}

/*
func Somaauxilioalimentacao(context *gin.Context) {
    var pdts []models.PDT
    config.DB.Find(&pdts)


    var iadvhs []models.Iadvh
    config.DB.Find(&iadvhs)


    somaPoradicionalnoturnoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao] += pdt.Adicional_noturno_20 + pdt.Adicional_noturno_30 + pdt.Adicional_noturno_20_valor + pdt.Media_adicional_noturno_salario_maternidade + pdt.Adicional_noturno_30_valor + pdt.Adicional_noturno_30_mes_anterior + pdt.Diferenca_adicional_noturno + pdt.Adicional_noturno_20_retroativo
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][iadvh.Unidade]; !ok {
            somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Adicional_noturno + iadvh.Adicional_noturno_mes_anterior
    }


    // Calcula a soma final e armazena em uma variável
 
//já foi



	somaPorauxilioalimentacaoUnidadeFuncaoPdt := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok :=  somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice] = make(map[string]map[string]float64)
        }
        if _, ok :=  somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice][pdt.Unidade]; !ok {
			somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice][pdt.Unidade] = make(map[string]float64)
        }
		somaPorauxilioalimentacaoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao] += pdt.Auxilio_alimentacao+pdt.Auxilio_alimentacao_retroativo
    }


    // Calcula a soma final e armazena em uma variável



//já foi




	somaPorpericulosidadeUnidadeFuncaoPdt  := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice][pdt.Unidade]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorpericulosidadeUnidadeFuncaoPdt [indice][pdt.Unidade][pdt.Funcao] += pdt.Periculosidade
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorpericulosidadeUnidadeFuncaoPdt [indice][iadvh.Unidade]; !ok {
            somaPorpericulosidadeUnidadeFuncaoPdt [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorpericulosidadeUnidadeFuncaoPdt [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Periculosidade
    }


    // Calcula a soma final e armazena em uma variável
  


	somaPorgratificacaoUnidadeFuncaoPdt   := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice][pdt.Unidade]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorgratificacaoUnidadeFuncaoPdt  [indice][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_funcao_retroativo+pdt.Gratificacao_funcao+pdt.Gratificacao_dedicacao_exclusiva+pdt.Gratificacao_produtividade+pdt.Gratificacao_funcao_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermage+pdt.Gratificacao_5_sobre_salario_base+pdt.Prorrogacao_media_deducao_exclusiva_salario_maternidade+pdt.Gratificacao_funcao_instrumentacao_cirurgico+pdt.Gratificacao_por_interiorizacao+pdt.Gratificacao+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_funcao_prorrogacao_maternidade+pdt.Gratificacao_funcao_supervisor_enfermagem_sl
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorgratificacaoUnidadeFuncaoPdt  [indice][iadvh.Unidade]; !ok {
            somaPorgratificacaoUnidadeFuncaoPdt  [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorgratificacaoUnidadeFuncaoPdt  [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Gratificacao+iadvh.Gratificacao_cargo_confianca+iadvh.Gratificacao_indenizatoria+iadvh.Gratificacao_funcao_instrumentacao_cirurgico+iadvh.Gratificacao_mes_anterior
    }


    // Calcula a soma final e armazena em uma variável




	
	somaPorhoraextraUnidadeFuncaoPdt    := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt   [indice]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt   [indice][pdt.Unidade]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorhoraextraUnidadeFuncaoPdt   [indice][pdt.Unidade][pdt.Funcao] += pdt.Media_horas_extras_prorrogacao_salario_maternidade+pdt.Hora_extra_50+pdt.Hora_extra_sobre_adicionais_50+pdt.Hora_extra_sobre_adicionais_60+pdt.Hora_extra_sobre_adicionais_100+pdt.Dsr_sobre_horas_extras+pdt.Dsr_sem_horas_extras_retroativas+pdt.Hora_extra_100+pdt.Hora_extra_retroativa_60+pdt.Hora_extra_50_valor+pdt.Hora_extra_100_mes_anterior+pdt.Hora_extra_50_retroativo+pdt.Hora_extra_100_valor+pdt.Hora_extra_60_valor
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt  [indice]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorhoraextraUnidadeFuncaoPdt   [indice][iadvh.Unidade]; !ok {
            somaPorhoraextraUnidadeFuncaoPdt   [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorhoraextraUnidadeFuncaoPdt   [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Hora_extra_100+iadvh.Hora_extra_50+iadvh.Hora_extra_60+iadvh.Hora_extra_60_mes_anterior+iadvh.Horas_extras_100_mes_anterior+iadvh.Horas_extras_50_mes_anterior
    }


    // Calcula a soma final e armazena em uma variável
    


	somaPorsalariobaseUnidadeFuncaoPdt    := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt   [indice]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt   [indice][pdt.Unidade]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorsalariobaseUnidadeFuncaoPdt   [indice][pdt.Unidade][pdt.Funcao] += pdt.Salario_base+pdt.Jovem_aprendiz+pdt.Diferenca_salarial+pdt.Devolucao_falta+pdt.Reembolso_faltas+pdt.Salario_substituicao+pdt.Salario_substituicao_mes_anterior+pdt.Estorno_desconto_vale_transporte+pdt.Dias_atestado+pdt.Prorrogacao_licenca_maternidade+pdt.Licenca_remunerada+pdt.Reembolso_educacional+pdt.Insuficiencia_saldo_provento+pdt.Saldo_salario+pdt.Saldo_salario_horista
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt    [indice]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorsalariobaseUnidadeFuncaoPdt   [indice][iadvh.Unidade]; !ok {
            somaPorsalariobaseUnidadeFuncaoPdt   [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorsalariobaseUnidadeFuncaoPdt   [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Atestado_medico+iadvh.Devolucao_desconto_indevido+iadvh.Devolucao_faltas+iadvh.Desconto_plano_saude_hapvida+iadvh.Desconto_plano_odontologico_integral+iadvh.Dias_trabalhados+iadvh.Diferenca_salarial+iadvh.Horas_noturnas_reduzidas+iadvh.Insuficiencia_saldo_proventos+iadvh.Licenca_paternidade+iadvh.Saldo_salario
    }


    // Calcula a soma final e armazena em uma variável
  


	somaPordeducaoUnidadeFuncaoPdt     := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice][pdt.Unidade]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPordeducaoUnidadeFuncaoPdt    [indice][pdt.Unidade][pdt.Funcao] += pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado 
    }


    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPordeducaoUnidadeFuncaoPdt    [indice][iadvh.Unidade]; !ok {
            somaPordeducaoUnidadeFuncaoPdt    [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPordeducaoUnidadeFuncaoPdt    [indice][iadvh.Unidade][iadvh.Funcao] += iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade
    }


    // Calcula a soma final e armazena em uma variável
   


	somaPorencargosUnidadeFuncaoPdt      := make(map[float64]map[string]map[string]float64)


    // Calcula a soma para pdts
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20 // Use um índice comum
        if _, ok := somaPorencargosUnidadeFuncaoPdt     [indice]; !ok {
            somaPorencargosUnidadeFuncaoPdt     [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorencargosUnidadeFuncaoPdt     [indice][pdt.Unidade]; !ok {
            somaPorencargosUnidadeFuncaoPdt   [indice][pdt.Unidade] = make(map[string]float64)
        }
        somaPorencargosUnidadeFuncaoPdt   [indice][pdt.Unidade][pdt.Funcao] += (pdt.Fgts_folha+pdt.Ferias+pdt.Ferias_nao_gozadas+pdt.Indenizacao_ferias_nao_gozadas+pdt.Ferias_sobre_aviso_indenizado+pdt.Aviso_previo_indenizado+pdt.Decimo_terceiro_sobre_aviso_indenizado+pdt.Ferias_vencidas+pdt.Ferias_proporcionais+pdt.Fgts_recisao+pdt.Fgts_decimo_terceiro_salario_recisao+pdt.Um_terco_ferias_recisao+pdt.Um_terco_ferias+pdt.Fgts_ferias) - (pdt.Gratificacao_produtor_salario_maternidade+pdt.Gratificacao_dedicacao_exclusiva_salario_maternidade+pdt.Gratificacao_interiorizacao_salario_maternidade+pdt.Salario_maternidade_jovem_aprendiz+pdt.Salario_maternidade+pdt.Media_horas_extras_salario_maternidade+pdt.Media_insalubridade_salario_maternidade+pdt.Gratificacao_sobre_maternidade_supervisor_enfermagem+pdt.Adicional_graduacao_15_tecnico_enfermagem_licenca_maternidade+pdt.Media_gratificacao_sobre_maternidade+pdt.Salario_familia+pdt.Salario_familia_indenizado )
    }



    // Calcula a soma para iadvhs
    for _, iadvh := range iadvhs {
        indice := iadvh.Adicional_noturno_mes_anterior // Use um índice comum
        if _, ok := somaPorencargosUnidadeFuncaoPdt    [indice]; !ok {
            somaPorencargosUnidadeFuncaoPdt     [indice] = make(map[string]map[string]float64)
        }
        if _, ok := somaPorencargosUnidadeFuncaoPdt     [indice][iadvh.Unidade]; !ok {
            somaPorencargosUnidadeFuncaoPdt     [indice][iadvh.Unidade] = make(map[string]float64)
        }
        somaPorencargosUnidadeFuncaoPdt     [indice][iadvh.Unidade][iadvh.Funcao] += (iadvh.Um_terco_ferias_mes+iadvh.Um_terco_ferias_proximo_mes+iadvh.Decimo_terceiro_indenizado+iadvh.Decimo_terceiro_licenca_maternidade+iadvh.Decimo_terceiro_recisao+iadvh.Adicional_um_terco_ferias+iadvh.Adicional_um_terco_ferias_proporcional_recisao+iadvh.Aviso_previo_indenizado+iadvh.Ferias_indenizadas+iadvh.Ferias_mes+iadvh.Ferias_proximo_mes+iadvh.Ferias_proporcionais+iadvh.Media_decimo_terceiro_rescisao+iadvh.Media_aviso_previo+iadvh.Media_ferias+iadvh.Media_ferias_proporcionais+iadvh.Multa_art_479_clt)- ( iadvh.Salario_maternidade+iadvh.Salario_familia+iadvh.Media_salario_maternidade)
    }


    // Calcula a soma final e armazena em uma variável
    var resultadoCalculo []map[string]interface{}
    for _, pdt := range pdts {
        indice := pdt.Adicional_noturno_20
        soma := somaPorencargosUnidadeFuncaoPdt    [indice][pdt.Unidade][pdt.Funcao]
        resultadoCalculo = append(resultadoCalculo, map[string]interface{}{
            "unidade": pdt.Unidade,
            "funcao":  pdt.Funcao,
            "adicional_noturno":somaPoradicionalnoturnoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao]idadeFuncaoPdt,
			"auxilio_alimentacao": somaPorpericulosidadeUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao],
			"periculosidade":somaPorpericulosidadeUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao],
			"gratificacao":somaPorgratificacaoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao] ,
			"hora_extra":somaPorhoraextraUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao], 
			"salario_folha": somaPorsalariobaseUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao],
			"deducao_folha": somaPordeducaoUnidadeFuncaoPdt[indice][pdt.Unidade][pdt.Funcao], 
			"soma_encargos": soma, 
			

        })
    }


    // Retorna o resultado sem persistir no banco de dados
    context.JSON(200, gin.H{
        "message": "Resultado do cálculo de Adicional_noturnofo",
        "result":  resultadoCalculo,
    })
}
*/