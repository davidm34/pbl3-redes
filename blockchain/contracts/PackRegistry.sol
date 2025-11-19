// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PackRegistry {
    // Variável de estado para armazenar o total de pacotes no ledger
    uint256 public totalPacks;

    // Evento para notificar clientes/servidores quando o estoque mudar
    event StockUpdated(uint256 newStock, string reason);

    // Define o estoque inicial ao fazer o deploy do contrato
    constructor(uint256 _initialStock) {
        totalPacks = _initialStock;
    }

    // Função para atualizar o estoque
    function setStock(uint256 _newStock, string memory _reason) public {
        totalPacks = _newStock;
        emit StockUpdated(_newStock, _reason);
    }

    // Função de leitura simples
    function getStock() public view returns (uint256) {
        return totalPacks;
    }
    
    // Função para decrementar um pacote (será usada na mecânica de jogo)
    function decrementStock() public {
        require(totalPacks > 0, "Estoque esgotado no Blockchain");
        totalPacks -= 1;
        emit StockUpdated(totalPacks, "Pacote aberto");
    }
}