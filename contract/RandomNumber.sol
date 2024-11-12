// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import {VRFConsumerBaseV2Plus} from "@chainlink/contracts/src/v0.8/vrf/dev/VRFConsumerBaseV2Plus.sol";
import {VRFV2PlusClient} from "@chainlink/contracts/src/v0.8/vrf/dev/libraries/VRFV2PlusClient.sol";

contract RandomNumber is VRFConsumerBaseV2Plus {
    error RandomNumber__RequestNotFound(uint256 requestId);              
    error RandomNumber__RequestAlreadyFulfilled(uint256 requestId);      
    error RandomNumber__RequestPending(uint256 requestId);               
    error RandomNumber__InvalidCoordinator(address coordinator);         
    error RandomNumber__InvalidSubscriptionId(uint64 subId);            
    error RandomNumber__InvalidCallbackGasLimit(uint32 gasLimit);       
    error RandomNumber__ZeroAddress();   

    struct RandomRequest {
        uint96 roundId;           // 减小到 uint96 节省 gas
        uint256[] randomNumbers;  // 存储随机数数组
        bool fulfilled;           // 是否已完成
        uint64 timestamp;         // 时间戳
    }

    // Polygon Amoy 网络配置
    address vrfCoordinator = 0x343300b5d84D444B2ADc9116FEF1bED02BE49Cf2;
    bytes32 s_keyHash = 0x816bedba8a50b294e5cbd47842baf240c2385f2eaf719edbd4f250a137a8c899;
    
    uint256 public s_subscriptionId;
    
    mapping(uint256 => RandomRequest) private s_randomRequests;
    uint96 private s_currentRound;
    
    event RequestedRandomness(
        uint256 indexed requestId, 
        uint96 indexed roundId,
        uint64 timestamp
    );
    
    event RandomnessFulfilled(
        uint256 indexed requestId, 
        uint96 indexed roundId, 
        uint256[] randomNumbers,
        uint64 timestamp
    );

    constructor(uint256 subscriptionId) VRFConsumerBaseV2Plus(vrfCoordinator) {
        s_subscriptionId = subscriptionId;
    }

    function requestRandomWords(
        uint32 numWords, 
        uint32 callbackGasLimit, 
        uint16 requestConfirmations
    ) public onlyOwner returns (uint256 s_requestId) {
        unchecked {
            s_currentRound++;
        }
        
        s_requestId = s_vrfCoordinator.requestRandomWords(
            VRFV2PlusClient.RandomWordsRequest({
                keyHash: s_keyHash,
                subId: s_subscriptionId,
                requestConfirmations: requestConfirmations,
                callbackGasLimit: callbackGasLimit,
                numWords: numWords,
                extraArgs: VRFV2PlusClient._argsToBytes(
                    VRFV2PlusClient.ExtraArgsV1({nativePayment: false})
                )
            })
        );
        
        s_randomRequests[s_requestId] = RandomRequest({
            roundId: s_currentRound,
            randomNumbers: new uint256[](0),
            fulfilled: false,
            timestamp: uint64(block.timestamp)
        });

        emit RequestedRandomness(
            s_requestId, 
            s_currentRound,
            uint64(block.timestamp)
        );
    }

    function fulfillRandomWords(
        uint256 requestId, 
        uint256[] calldata randomWords
    ) internal override {
        RandomRequest storage request = s_randomRequests[requestId];

        if (request.roundId == 0) revert RandomNumber__RequestNotFound(requestId);
        if (request.fulfilled) revert RandomNumber__RequestAlreadyFulfilled(requestId);

        request.randomNumbers = randomWords;
        request.fulfilled = true;
        request.timestamp = uint64(block.timestamp);

        emit RandomnessFulfilled(
            requestId,
            request.roundId,
            randomWords,
            uint64(block.timestamp)
        );
    }

    function getRandomRequest(
        uint256 requestId
    ) external view returns (RandomRequest memory) {
        RandomRequest memory request = s_randomRequests[requestId];
        if (request.roundId == 0) revert RandomNumber__RequestNotFound(requestId);
        return request;
    }

    function getLatestRandomNumber(
        uint256 requestId
    ) external view returns (uint256[] memory) {
        RandomRequest memory request = s_randomRequests[requestId];
        if (request.roundId == 0) revert RandomNumber__RequestNotFound(requestId);
        if (!request.fulfilled) revert RandomNumber__RequestPending(requestId);
        return request.randomNumbers;
    }

    function getCurrentRound() external view returns (uint96) {
        return s_currentRound;
    }
}