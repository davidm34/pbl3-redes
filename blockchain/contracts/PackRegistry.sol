// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PackRegistry {
    // --- Parte 1: Estoque de Pacotes (Fase 2) ---
    uint256 public totalPacks;
    event StockUpdated(uint256 newStock, string reason);

    constructor(uint256 _initialStock) {
        totalPacks = _initialStock;
    }

    function decrementStock() public {
        require(totalPacks > 0, "Estoque esgotado no Blockchain");
        totalPacks -= 1;
        emit StockUpdated(totalPacks, "Pacote aberto");
    }

    function getStock() public view returns (uint256) {
        return totalPacks;
    }

    // --- Parte 2: Registo de Partidas (Fase 3) ---
    
    struct Match {
        string matchId;
        string winnerId;
        string loserId;
        uint256 timestamp;
    }

    // Lista pública de todas as partidas
    Match[] public matches;

    // Evento para facilitar a leitura por clientes externos
    event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp);

    function recordMatch(string memory _matchId, string memory _winnerId, string memory _loserId) public {
        matches.push(Match(_matchId, _winnerId, _loserId, block.timestamp));
        emit MatchRecorded(_matchId, _winnerId, _loserId, block.timestamp);
    }

    function getMatchCount() public view returns (uint256) {
        return matches.length;
    }
}