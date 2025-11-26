
-----

# 🐜 ACO Solver (Ant Colony Optimization)

Uma implementação robusta, modular e de alta performance do algoritmo **Otimização por Colônia de Formigas (ACO)** escrita em **Go** para resolver o **Problema do Caixeiro Viajante (TSP)**.

Este projeto utiliza **concorrência (Goroutines)** para simular formigas em paralelo e gera visualizações detalhadas para análise do comportamento do algoritmo.

-----

## 🚀 Funcionalidades

  * 🧠 **Algoritmo ACO Clássico**: Baseado em feromônios e heurísticas de visibilidade.
  * ⚡ **Otimização com KNN**: Usa *K-Nearest Neighbors* para reduzir o espaço de busca.
  * 🔥 **Execução Paralela**: Caminhos calculados simultaneamente para máxima velocidade.
  * 📊 **Visualização Rica**: Gera gráficos automáticos (`.png`) ao final da execução:
      * Convergência de custo.
      * Melhor rota no mapa.
      * Heatmaps de feromônios.
      * Estatísticas de tempo.
  * ⚙️ **Configuração Flexível**: Suporte a CLI, JSON e execução em Batch (Grid Search).

-----

## 📦 Instalação

Certifique-se de ter o [Go instalado](https://go.dev/dl/).

1.  **Clone o repositório:**

    ```bash
    git clone https://github.com/SRwasabi/ACO_att.git
    cd ACO_att
    ```

2.  **Baixe as dependências:**

    ```bash
    go mod tidy
    ```

-----

## 🛠️ Como Usar

### 1. Execução Rápida (CLI)

Rode diretamente pelo terminal passando parâmetros:

```bash
go run . -file coordinates/wi29.tsp -ants 100 -iter 500
```

### 2. Experimentos em Batch (JSON)

Para testar múltiplos parâmetros de uma vez (Grid Search), crie um arquivo `config.json`:

```json
{
    "experiment_name": "Teste_Grid",
    "input_file": ["coordinates/wi29.tsp"],
    "num_ants": [50, 100],
    "iterations": [200],
    "alpha": [1.0],
    "beta": [2.0, 5.0],
    "evaporation": [0.1],
    "q": [100.0],
    "knn_size": [20],
    "use_knn": [true],
    "roulette_selection": [true],
    "seed": 0
}
```

Execute com:

```bash
go run . -config config.json
```

----- 

### 3. Modo Híbrido (CLI + Config)

É possível combinar um arquivo de configuração com flags na linha de comando. Regras principais:

- Flags em CLI têm precedência sobre valores correspondentes em `config.json` (sobrescrevem).
- Flags que recebem uma lista (ex.: -iter "100,200") são interpretadas como múltiplos valores e participam do Grid Search (produto cartesiano com listas do config).
- Flags com um único valor (ex.: -ants 100) substituem aquele parâmetro em todas as execuções geradas pelo config.

Exemplos:

- Sobrescrever um valor simples:
```bash
go run . -config config.json -ants 100
```
(Irá usar `ants=100` para todas as execuções definidas pelo config.)

- Forçar múltiplos valores via CLI (expande o grid):
```bash
go run . -config config.json -ants 100 -iter "100,200"
```
(Irá combinar o `config.json` com `ants=100` e `iterations` em {100,200}, gerando execuções para cada combinação.)

Observação: sempre coloque listas entre aspas para evitar interpretação da shell (ex.: "100,200").  

----- 

## 🌍 Datasets (.tsp)

O projeto já inclui alguns exemplos na pasta `coordinates/`. Para testar com mapas reais de países (como Djibuti, Catar, Uruguai, etc.), você pode baixar arquivos `.tsp` compatíveis no link abaixo:

🔗 **[National TSP Instances (University of Waterloo)](https://www.math.uwaterloo.ca/tsp/world/countries.html)**

> **Dica:** Baixe o arquivo, salve na pasta `coordinates/` e aponte o caminho na execução (ex: `-file coordinates/uy734.tsp`).

-----

## 📊 Resultados Gerados

Após a execução, uma pasta `NomeDoExperimento_Results/` será criada com:

| Arquivo | Descrição |
| :--- | :--- |
| `best_path.png` | 🗺️ Desenho da melhor rota encontrada. |
| `convergence.png` | 📉 Gráfico da evolução do custo por iteração. |
| `costsStats.png` | 📊 Comparativo (Melhor vs Médio vs Pior). |
| `heatmap_*.png` | 🔥 Intensidade das trilhas de feromônio. |
| `timing.png` | ⏱️ Tempo gasto em cada fase do algoritmo. |

-----

## 📂 Estrutura do Projeto

```text
.
├── coordinates/   # Arquivos de dados (.tsp)
├── pkg/
│   ├── aco/       # Lógica da Colônia e Formigas
│   ├── config/    # Carregamento de JSON e Flags
│   ├── graph/     # Grafo, Cidades e KNN
│   └── plotter/   # Geração de gráficos e imagens
├── main.go        # Ponto de entrada principal
├── go.mod         # Dependências do G
└── README.md      # Este arquivo
```

-----

## 👥 Contribuidores

Projeto desenvolvido com ❤️ por:

  * **[SRwasabi](https://github.com/SRwasabi)**
  * **[Felipe Camarano](https://github.com/FelipeCamarano)**
  * **[Felipe Camarano](https://github.com/FelipeCamaranoInpulso)**

-----
