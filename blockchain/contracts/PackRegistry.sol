// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract PackRegistry {
    uint256 public totalPacks;
    struct Match { string matchId; string winnerId; string loserId; uint256 timestamp; }
    Match[] public matches;
    event StockUpdated(uint256 newStock, string reason);
    event MatchRecorded(string matchId, string winner, string loser, uint256 timestamp);
    
    // Mapeia ID da Carta -> Dono Atual
    mapping(string => string) public cardOwner;
    
    // Mapeia Dono -> Lista de IDs de Cartas (Inventário)
    mapping(string => string[]) public ownerCards;

    event CardAssigned(string cardId, string ownerId, uint256 timestamp);
    event CardTransferred(string cardId, string fromId, string toId, uint256 timestamp);

    constructor(uint256 _initialStock) {
        totalPacks = _initialStock;
    }

    function decrementStock() public {
        require(totalPacks > 0, "Estoque esgotado");
        totalPacks -= 1;
        emit StockUpdated(totalPacks, "Pacote aberto");
    }
    function getStock() public view returns (uint256) { return totalPacks; }
    function recordMatch(string memory _matchId, string memory _winnerId, string memory _loserId) public {
        matches.push(Match(_matchId, _winnerId, _loserId, block.timestamp));
        emit MatchRecorded(_matchId, _winnerId, _loserId, block.timestamp);
    }
    function getMatchCount() public view returns (uint256) { return matches.length; }

    function assignCards(string memory _playerId, string[] memory _cardIds) public {
        for (uint i = 0; i < _cardIds.length; i++) {
            string memory cardId = _cardIds[i];
            require(bytes(cardOwner[cardId]).length == 0, "Carta ja tem dono");
            
            cardOwner[cardId] = _playerId;
            ownerCards[_playerId].push(cardId); // Adiciona à lista do dono
            
            emit CardAssigned(cardId, _playerId, block.timestamp);
        }
    }

    function transferCard(string memory _fromId, string memory _toId, string memory _cardId) public {
        require(keccak256(bytes(cardOwner[_cardId])) == keccak256(bytes(_fromId)), "Nao e o dono");
        
        // 1. Atualiza o dono
        cardOwner[_cardId] = _toId;

        // 2. Remove da lista do antigo dono
        removeCardFromList(_fromId, _cardId);

        // 3. Adiciona à lista do novo dono
        ownerCards[_toId].push(_cardId);

        emit CardTransferred(_cardId, _fromId, _toId, block.timestamp);
    }

    // Função auxiliar para remover carta da lista (Swap & Pop)
    function removeCardFromList(string memory _ownerId, string memory _cardId) internal {
        string[] storage cards = ownerCards[_ownerId];
        for (uint i = 0; i < cards.length; i++) {
            if (keccak256(bytes(cards[i])) == keccak256(bytes(_cardId))) {
                // Substitui o elemento a remover pelo último da lista
                cards[i] = cards[cards.length - 1];
                // Remove o último elemento
                cards.pop();
                break;
            }
        }
    }

    // NOVO: Função para o usuário visualizar suas cartas
    function getUserCards(string memory _ownerId) public view returns (string[] memory) {
        return ownerCards[_ownerId];
    }
}