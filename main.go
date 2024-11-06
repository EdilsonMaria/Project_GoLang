//Servidor HTTP simples escrito em Go
package main

//Importando os packager nescessarios (onverter os dados em Json, configurar o servidor HTTP e formatar datas e horários.)
import (
    "encoding/json"
    "net/http"
    "time"
)

//Define uma estrutura chamada Response, que representa a resposta enviada ao cliente.
type Response struct {
    Nome    string `json:"nome"`
    Horario string `json:"horario"`
}

//Handler é chamada sempre que uma requisição é feita ao servidor.
func handler(w http.ResponseWriter, r *http.Request) {

    //Definição do fuso horário com base na localização
    location, err := time.LoadLocation("America/Sao_Paulo")
    if err != nil {
        http.Error(w, "Erro ao carregar o fuso horário", http.StatusInternalServerError)
        return
    }

    //Criação da resposta
    response := Response{
        Nome:    "Projeto Korp",
        Horario: time.Now().In(location).Format(time.RFC3339),
    }
    
	//Configurar e enviar a resposta (cabeçalho)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

//Define que, quando o servidor receber uma requisição no caminho raiz (/), ele chamará a função handler
func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}