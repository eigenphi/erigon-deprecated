// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Block.proto

package io.eigenphi.erigon.protos;

public interface InternalTxnOrBuilder extends
    // @@protoc_insertion_point(interface_extends:erigon.InternalTxn)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <code>string transactionHash = 1;</code>
   * @return The transactionHash.
   */
  java.lang.String getTransactionHash();
  /**
   * <code>string transactionHash = 1;</code>
   * @return The bytes for transactionHash.
   */
  com.google.protobuf.ByteString
      getTransactionHashBytes();

  /**
   * <code>string from = 2;</code>
   * @return The from.
   */
  java.lang.String getFrom();
  /**
   * <code>string from = 2;</code>
   * @return The bytes for from.
   */
  com.google.protobuf.ByteString
      getFromBytes();

  /**
   * <code>string to = 3;</code>
   * @return The to.
   */
  java.lang.String getTo();
  /**
   * <code>string to = 3;</code>
   * @return The bytes for to.
   */
  com.google.protobuf.ByteString
      getToBytes();

  /**
   * <code>string value = 4;</code>
   * @return The value.
   */
  java.lang.String getValue();
  /**
   * <code>string value = 4;</code>
   * @return The bytes for value.
   */
  com.google.protobuf.ByteString
      getValueBytes();
}