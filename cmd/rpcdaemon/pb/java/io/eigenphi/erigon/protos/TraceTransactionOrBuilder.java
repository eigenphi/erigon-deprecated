// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Block.proto

package io.eigenphi.erigon.protos;

public interface TraceTransactionOrBuilder extends
    // @@protoc_insertion_point(interface_extends:erigon.TraceTransaction)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>int64 blockNumber = 1;</code>
   * @return The blockNumber.
   */
  long getBlockNumber();

  /**
   * <code>string transactionHash = 2;</code>
   * @return The transactionHash.
   */
  java.lang.String getTransactionHash();
  /**
   * <code>string transactionHash = 2;</code>
   * @return The bytes for transactionHash.
   */
  com.google.protobuf.ByteString
      getTransactionHashBytes();

  /**
   * <code>int32 transactionIndex = 3;</code>
   * @return The transactionIndex.
   */
  int getTransactionIndex();

  /**
   * <code>string fromAddress = 4;</code>
   * @return The fromAddress.
   */
  java.lang.String getFromAddress();
  /**
   * <code>string fromAddress = 4;</code>
   * @return The bytes for fromAddress.
   */
  com.google.protobuf.ByteString
      getFromAddressBytes();

  /**
   * <code>string toAddress = 5;</code>
   * @return The toAddress.
   */
  java.lang.String getToAddress();
  /**
   * <code>string toAddress = 5;</code>
   * @return The bytes for toAddress.
   */
  com.google.protobuf.ByteString
      getToAddressBytes();

  /**
   * <code>int64 gasPrice = 6;</code>
   * @return The gasPrice.
   */
  long getGasPrice();

  /**
   * <code>string input = 7;</code>
   * @return The input.
   */
  java.lang.String getInput();
  /**
   * <code>string input = 7;</code>
   * @return The bytes for input.
   */
  com.google.protobuf.ByteString
      getInputBytes();

  /**
   * <code>int64 nonce = 8;</code>
   * @return The nonce.
   */
  long getNonce();

  /**
   * <code>string transactionValue = 9;</code>
   * @return The transactionValue.
   */
  java.lang.String getTransactionValue();
  /**
   * <code>string transactionValue = 9;</code>
   * @return The bytes for transactionValue.
   */
  com.google.protobuf.ByteString
      getTransactionValueBytes();

  /**
   * <code>.erigon.stackFrame stack = 10;</code>
   * @return Whether the stack field is set.
   */
  boolean hasStack();
  /**
   * <code>.erigon.stackFrame stack = 10;</code>
   * @return The stack.
   */
  io.eigenphi.erigon.protos.stackFrame getStack();
  /**
   * <code>.erigon.stackFrame stack = 10;</code>
   */
  io.eigenphi.erigon.protos.stackFrameOrBuilder getStackOrBuilder();
}
