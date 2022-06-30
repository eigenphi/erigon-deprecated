// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Block.proto

package io.eigenphi.erigon.protos;

public interface EventLogOrBuilder extends
    // @@protoc_insertion_point(interface_extends:erigon.EventLog)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string chain = 1;</code>
   * @return The chain.
   */
  java.lang.String getChain();
  /**
   * <code>string chain = 1;</code>
   * @return The bytes for chain.
   */
  com.google.protobuf.ByteString
      getChainBytes();

  /**
   * <code>string address = 2;</code>
   * @return The address.
   */
  java.lang.String getAddress();
  /**
   * <code>string address = 2;</code>
   * @return The bytes for address.
   */
  com.google.protobuf.ByteString
      getAddressBytes();

  /**
   * <code>string topic0 = 3;</code>
   * @return The topic0.
   */
  java.lang.String getTopic0();
  /**
   * <code>string topic0 = 3;</code>
   * @return The bytes for topic0.
   */
  com.google.protobuf.ByteString
      getTopic0Bytes();

  /**
   * <code>string topic1 = 4;</code>
   * @return The topic1.
   */
  java.lang.String getTopic1();
  /**
   * <code>string topic1 = 4;</code>
   * @return The bytes for topic1.
   */
  com.google.protobuf.ByteString
      getTopic1Bytes();

  /**
   * <code>string topic2 = 5;</code>
   * @return The topic2.
   */
  java.lang.String getTopic2();
  /**
   * <code>string topic2 = 5;</code>
   * @return The bytes for topic2.
   */
  com.google.protobuf.ByteString
      getTopic2Bytes();

  /**
   * <code>string topic3 = 6;</code>
   * @return The topic3.
   */
  java.lang.String getTopic3();
  /**
   * <code>string topic3 = 6;</code>
   * @return The bytes for topic3.
   */
  com.google.protobuf.ByteString
      getTopic3Bytes();

  /**
   * <code>string data = 7;</code>
   * @return The data.
   */
  java.lang.String getData();
  /**
   * <code>string data = 7;</code>
   * @return The bytes for data.
   */
  com.google.protobuf.ByteString
      getDataBytes();

  /**
   * <code>int64 blockNumber = 8;</code>
   * @return The blockNumber.
   */
  long getBlockNumber();

  /**
   * <code>int64 blockTimestamp = 9;</code>
   * @return The blockTimestamp.
   */
  long getBlockTimestamp();

  /**
   * <code>string transactionHash = 10;</code>
   * @return The transactionHash.
   */
  java.lang.String getTransactionHash();
  /**
   * <code>string transactionHash = 10;</code>
   * @return The bytes for transactionHash.
   */
  com.google.protobuf.ByteString
      getTransactionHashBytes();

  /**
   * <code>int32 transactionIndex = 11;</code>
   * @return The transactionIndex.
   */
  int getTransactionIndex();

  /**
   * <code>string blockHash = 12;</code>
   * @return The blockHash.
   */
  java.lang.String getBlockHash();
  /**
   * <code>string blockHash = 12;</code>
   * @return The bytes for blockHash.
   */
  com.google.protobuf.ByteString
      getBlockHashBytes();

  /**
   * <code>int32 logIndex = 13;</code>
   * @return The logIndex.
   */
  int getLogIndex();

  /**
   * <code>bool removed = 14;</code>
   * @return The removed.
   */
  boolean getRemoved();

  /**
   * <code>string senderInfo = 15;</code>
   * @return The senderInfo.
   */
  java.lang.String getSenderInfo();
  /**
   * <code>string senderInfo = 15;</code>
   * @return The bytes for senderInfo.
   */
  com.google.protobuf.ByteString
      getSenderInfoBytes();
}
