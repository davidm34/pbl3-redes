// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PackRegistry {
    // --- Fase 2: Estoque ---
    uint256 public totalPacks;
    event StockUpdated(uint256 newStock, string reason);

    // --- Fase 3: Histórico de Partidas ---
    struct Match {
        string matchId;
        string winnerId;
        string loserId;
        uint256 timestamp;
    }
    Match[] public matches;
    event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp);

    // --- FASE 4: PROPRIEDADE DE CARTAS ---
    
    // Mapeia ID da Carta (ex: "c_001_abc123") -> ID do Jogador (ex: "player_1")
    // Usamos string para IDs porque o seu sistema atual usa strings.
    mapping(string => string) public cardOwner;

    // Eventos para auditoria
    event CardAssigned(string cardId, string ownerId, uint256 timestamp);
    event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp);

    constructor(uint256 _initialStock) {
        totalPacks = _initialStock;
    }

    // --- Funções de Estoque ---
    function decrementStock() public {
        require(totalPacks > 0, "Estoque esgotado no Blockchain");
        totalPacks -= 1;
        emit StockUpdated(totalPacks, "Pacote aberto");
    }

    function getStock() public view returns (uint256) {
        return totalPacks;
    }

    // --- Funções de Partidas ---
    function recordMatch(string memory _matchId, string memory _winnerId, string memory _loserId) public {
        matches.push(Match(_matchId, _winnerId, _loserId, block.timestamp));
        emit MatchRecorded(_matchId, _winnerId, _loserId, block.timestamp);
    }

    function getMatchCount() public view returns (uint256) {
        return matches.length;
    }

    // --- NOVAS FUNÇÕES FASE 4 ---

    // Atribui uma lista de cartas a um jogador (chamado após abrir pacote)
    function assignCards(string memory _playerId, string[] memory _cardIds) public {
        for (uint i = 0; i < _cardIds.length; i++) {
            string memory cardId = _cardIds[i];
            // Garante que a carta ainda não tem dono (evita duplicatas/fraude)
            require(bytes(cardOwner[cardId]).length == 0, "Carta ja tem dono");
            
            cardOwner[cardId] = _playerId;
            emit CardAssigned(cardId, _playerId, block.timestamp);
        }
    }

    // Transfere a posse de uma carta de um jogador para outro
    function transferCard(string memory _fromId, string memory _toId, string memory _cardId) public {
        // Verifica se o remetente é realmente o dono atual
        // Nota: Comparação de strings em Solidity requer hash (keccak256)
        require(keccak256(bytes(cardOwner[_cardId])) == keccak256(bytes(_fromId)), "Remetente nao possui esta carta");
        
        cardOwner[_cardId] = _toId;
        emit CardTransferred(_cardId, _fromId, _toId, block.timestamp);
    }

    // Função auxiliar para verificar dono (opcional, para testes)
    function getCardOwner(string memory _cardId) public view returns (string memory) {
        return cardOwner[_cardId];
    }
}