// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Block.proto

package io.eigenphi.erigon.protos;

public interface stackFrameOrBuilder extends
    // @@protoc_insertion_point(interface_extends:erigon.stackFrame)
    com.google.protobuf.MessageOrBuilder {

  /**
   * <pre>
   * OP code
   * </pre>
   *
   * <code>string type = 1;</code>
   * @return The type.
   */
  java.lang.String getType();
  /**
   * <pre>
   * OP code
   * </pre>
   *
   * <code>string type = 1;</code>
   * @return The bytes for type.
   */
  com.google.protobuf.ByteString
      getTypeBytes();

  /**
   * <pre>
   * empty or "Transfer" or "Internal-Transfer"
   * </pre>
   *
   * <code>string label = 2;</code>
   * @return The label.
   */
  java.lang.String getLabel();
  /**
   * <pre>
   * empty or "Transfer" or "Internal-Transfer"
   * </pre>
   *
   * <code>string label = 2;</code>
   * @return The bytes for label.
   */
  com.google.protobuf.ByteString
      getLabelBytes();

  /**
   * <code>string from = 3;</code>
   * @return The from.
   */
  java.lang.String getFrom();
  /**
   * <code>string from = 3;</code>
   * @return The bytes for from.
   */
  com.google.protobuf.ByteString
      getFromBytes();

  /**
   * <code>string to = 4;</code>
   * @return The to.
   */
  java.lang.String getTo();
  /**
   * <code>string to = 4;</code>
   * @return The bytes for to.
   */
  com.google.protobuf.ByteString
      getToBytes();

  /**
   * <code>string contractCreated = 5;</code>
   * @return The contractCreated.
   */
  java.lang.String getContractCreated();
  /**
   * <code>string contractCreated = 5;</code>
   * @return The bytes for contractCreated.
   */
  com.google.protobuf.ByteString
      getContractCreatedBytes();

  /**
   * <code>string value = 6;</code>
   * @return The value.
   */
  java.lang.String getValue();
  /**
   * <code>string value = 6;</code>
   * @return The bytes for value.
   */
  com.google.protobuf.ByteString
      getValueBytes();

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
   * <code>string error = 8;</code>
   * @return The error.
   */
  java.lang.String getError();
  /**
   * <code>string error = 8;</code>
   * @return The bytes for error.
   */
  com.google.protobuf.ByteString
      getErrorBytes();

  /**
   * <code>repeated .erigon.stackFrame calls = 9;</code>
   */
  java.util.List<io.eigenphi.erigon.protos.stackFrame> 
      getCallsList();
  /**
   * <code>repeated .erigon.stackFrame calls = 9;</code>
   */
  io.eigenphi.erigon.protos.stackFrame getCalls(int index);
  /**
   * <code>repeated .erigon.stackFrame calls = 9;</code>
   */
  int getCallsCount();
  /**
   * <code>repeated .erigon.stackFrame calls = 9;</code>
   */
  java.util.List<? extends io.eigenphi.erigon.protos.stackFrameOrBuilder> 
      getCallsOrBuilderList();
  /**
   * <code>repeated .erigon.stackFrame calls = 9;</code>
   */
  io.eigenphi.erigon.protos.stackFrameOrBuilder getCallsOrBuilder(
      int index);

  /**
   * <code>string fourBytes = 10;</code>
   * @return The fourBytes.
   */
  java.lang.String getFourBytes();
  /**
   * <code>string fourBytes = 10;</code>
   * @return The bytes for fourBytes.
   */
  com.google.protobuf.ByteString
      getFourBytesBytes();
}
