// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Block.proto

package io.eigenphi.erigon.protos;

/**
 * Protobuf type {@code erigon.Transaction}
 */
public final class Transaction extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:erigon.Transaction)
    TransactionOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Transaction.newBuilder() to construct.
  private Transaction(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Transaction() {
    transactionHash_ = "";
    fromAddress_ = "";
    toAddress_ = "";
    transactionValue_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Transaction();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Transaction_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Transaction_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.eigenphi.erigon.protos.Transaction.class, io.eigenphi.erigon.protos.Transaction.Builder.class);
  }

  public static final int TRANSACTIONHASH_FIELD_NUMBER = 1;
  private volatile java.lang.Object transactionHash_;
  /**
   * <code>string transactionHash = 1;</code>
   * @return The transactionHash.
   */
  @java.lang.Override
  public java.lang.String getTransactionHash() {
    java.lang.Object ref = transactionHash_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      transactionHash_ = s;
      return s;
    }
  }
  /**
   * <code>string transactionHash = 1;</code>
   * @return The bytes for transactionHash.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getTransactionHashBytes() {
    java.lang.Object ref = transactionHash_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      transactionHash_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int TRANSACTIONINDEX_FIELD_NUMBER = 2;
  private int transactionIndex_;
  /**
   * <code>int32 transactionIndex = 2;</code>
   * @return The transactionIndex.
   */
  @java.lang.Override
  public int getTransactionIndex() {
    return transactionIndex_;
  }

  public static final int BLOCKNUMBER_FIELD_NUMBER = 3;
  private long blockNumber_;
  /**
   * <code>int64 blockNumber = 3;</code>
   * @return The blockNumber.
   */
  @java.lang.Override
  public long getBlockNumber() {
    return blockNumber_;
  }

  public static final int FROMADDRESS_FIELD_NUMBER = 4;
  private volatile java.lang.Object fromAddress_;
  /**
   * <code>string fromAddress = 4;</code>
   * @return The fromAddress.
   */
  @java.lang.Override
  public java.lang.String getFromAddress() {
    java.lang.Object ref = fromAddress_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      fromAddress_ = s;
      return s;
    }
  }
  /**
   * <code>string fromAddress = 4;</code>
   * @return The bytes for fromAddress.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getFromAddressBytes() {
    java.lang.Object ref = fromAddress_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      fromAddress_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int TOADDRESS_FIELD_NUMBER = 5;
  private volatile java.lang.Object toAddress_;
  /**
   * <code>string toAddress = 5;</code>
   * @return The toAddress.
   */
  @java.lang.Override
  public java.lang.String getToAddress() {
    java.lang.Object ref = toAddress_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      toAddress_ = s;
      return s;
    }
  }
  /**
   * <code>string toAddress = 5;</code>
   * @return The bytes for toAddress.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getToAddressBytes() {
    java.lang.Object ref = toAddress_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      toAddress_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int NONCE_FIELD_NUMBER = 6;
  private long nonce_;
  /**
   * <code>int64 nonce = 6;</code>
   * @return The nonce.
   */
  @java.lang.Override
  public long getNonce() {
    return nonce_;
  }

  public static final int TRANSACTIONVALUE_FIELD_NUMBER = 7;
  private volatile java.lang.Object transactionValue_;
  /**
   * <code>string transactionValue = 7;</code>
   * @return The transactionValue.
   */
  @java.lang.Override
  public java.lang.String getTransactionValue() {
    java.lang.Object ref = transactionValue_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      transactionValue_ = s;
      return s;
    }
  }
  /**
   * <code>string transactionValue = 7;</code>
   * @return The bytes for transactionValue.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getTransactionValueBytes() {
    java.lang.Object ref = transactionValue_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      transactionValue_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int BLOCKTIMESTAMP_FIELD_NUMBER = 8;
  private long blockTimestamp_;
  /**
   * <code>int64 blockTimestamp = 8;</code>
   * @return The blockTimestamp.
   */
  @java.lang.Override
  public long getBlockTimestamp() {
    return blockTimestamp_;
  }

  private byte memoizedIsInitialized = -1;
  @java.lang.Override
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  @java.lang.Override
  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(transactionHash_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, transactionHash_);
    }
    if (transactionIndex_ != 0) {
      output.writeInt32(2, transactionIndex_);
    }
    if (blockNumber_ != 0L) {
      output.writeInt64(3, blockNumber_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(fromAddress_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 4, fromAddress_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(toAddress_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 5, toAddress_);
    }
    if (nonce_ != 0L) {
      output.writeInt64(6, nonce_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(transactionValue_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 7, transactionValue_);
    }
    if (blockTimestamp_ != 0L) {
      output.writeInt64(8, blockTimestamp_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(transactionHash_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, transactionHash_);
    }
    if (transactionIndex_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(2, transactionIndex_);
    }
    if (blockNumber_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(3, blockNumber_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(fromAddress_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, fromAddress_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(toAddress_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(5, toAddress_);
    }
    if (nonce_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(6, nonce_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(transactionValue_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(7, transactionValue_);
    }
    if (blockTimestamp_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(8, blockTimestamp_);
    }
    size += getUnknownFields().getSerializedSize();
    memoizedSize = size;
    return size;
  }

  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof io.eigenphi.erigon.protos.Transaction)) {
      return super.equals(obj);
    }
    io.eigenphi.erigon.protos.Transaction other = (io.eigenphi.erigon.protos.Transaction) obj;

    if (!getTransactionHash()
        .equals(other.getTransactionHash())) return false;
    if (getTransactionIndex()
        != other.getTransactionIndex()) return false;
    if (getBlockNumber()
        != other.getBlockNumber()) return false;
    if (!getFromAddress()
        .equals(other.getFromAddress())) return false;
    if (!getToAddress()
        .equals(other.getToAddress())) return false;
    if (getNonce()
        != other.getNonce()) return false;
    if (!getTransactionValue()
        .equals(other.getTransactionValue())) return false;
    if (getBlockTimestamp()
        != other.getBlockTimestamp()) return false;
    if (!getUnknownFields().equals(other.getUnknownFields())) return false;
    return true;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    hash = (37 * hash) + TRANSACTIONHASH_FIELD_NUMBER;
    hash = (53 * hash) + getTransactionHash().hashCode();
    hash = (37 * hash) + TRANSACTIONINDEX_FIELD_NUMBER;
    hash = (53 * hash) + getTransactionIndex();
    hash = (37 * hash) + BLOCKNUMBER_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getBlockNumber());
    hash = (37 * hash) + FROMADDRESS_FIELD_NUMBER;
    hash = (53 * hash) + getFromAddress().hashCode();
    hash = (37 * hash) + TOADDRESS_FIELD_NUMBER;
    hash = (53 * hash) + getToAddress().hashCode();
    hash = (37 * hash) + NONCE_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getNonce());
    hash = (37 * hash) + TRANSACTIONVALUE_FIELD_NUMBER;
    hash = (53 * hash) + getTransactionValue().hashCode();
    hash = (37 * hash) + BLOCKTIMESTAMP_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getBlockTimestamp());
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Transaction parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.eigenphi.erigon.protos.Transaction parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.eigenphi.erigon.protos.Transaction parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  @java.lang.Override
  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(io.eigenphi.erigon.protos.Transaction prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  @java.lang.Override
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code erigon.Transaction}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:erigon.Transaction)
      io.eigenphi.erigon.protos.TransactionOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Transaction_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Transaction_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.eigenphi.erigon.protos.Transaction.class, io.eigenphi.erigon.protos.Transaction.Builder.class);
    }

    // Construct using io.eigenphi.erigon.protos.Transaction.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      transactionHash_ = "";

      transactionIndex_ = 0;

      blockNumber_ = 0L;

      fromAddress_ = "";

      toAddress_ = "";

      nonce_ = 0L;

      transactionValue_ = "";

      blockTimestamp_ = 0L;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Transaction_descriptor;
    }

    @java.lang.Override
    public io.eigenphi.erigon.protos.Transaction getDefaultInstanceForType() {
      return io.eigenphi.erigon.protos.Transaction.getDefaultInstance();
    }

    @java.lang.Override
    public io.eigenphi.erigon.protos.Transaction build() {
      io.eigenphi.erigon.protos.Transaction result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.eigenphi.erigon.protos.Transaction buildPartial() {
      io.eigenphi.erigon.protos.Transaction result = new io.eigenphi.erigon.protos.Transaction(this);
      result.transactionHash_ = transactionHash_;
      result.transactionIndex_ = transactionIndex_;
      result.blockNumber_ = blockNumber_;
      result.fromAddress_ = fromAddress_;
      result.toAddress_ = toAddress_;
      result.nonce_ = nonce_;
      result.transactionValue_ = transactionValue_;
      result.blockTimestamp_ = blockTimestamp_;
      onBuilt();
      return result;
    }

    @java.lang.Override
    public Builder clone() {
      return super.clone();
    }
    @java.lang.Override
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.setField(field, value);
    }
    @java.lang.Override
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return super.clearField(field);
    }
    @java.lang.Override
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return super.clearOneof(oneof);
    }
    @java.lang.Override
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, java.lang.Object value) {
      return super.setRepeatedField(field, index, value);
    }
    @java.lang.Override
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        java.lang.Object value) {
      return super.addRepeatedField(field, value);
    }
    @java.lang.Override
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof io.eigenphi.erigon.protos.Transaction) {
        return mergeFrom((io.eigenphi.erigon.protos.Transaction)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.eigenphi.erigon.protos.Transaction other) {
      if (other == io.eigenphi.erigon.protos.Transaction.getDefaultInstance()) return this;
      if (!other.getTransactionHash().isEmpty()) {
        transactionHash_ = other.transactionHash_;
        onChanged();
      }
      if (other.getTransactionIndex() != 0) {
        setTransactionIndex(other.getTransactionIndex());
      }
      if (other.getBlockNumber() != 0L) {
        setBlockNumber(other.getBlockNumber());
      }
      if (!other.getFromAddress().isEmpty()) {
        fromAddress_ = other.fromAddress_;
        onChanged();
      }
      if (!other.getToAddress().isEmpty()) {
        toAddress_ = other.toAddress_;
        onChanged();
      }
      if (other.getNonce() != 0L) {
        setNonce(other.getNonce());
      }
      if (!other.getTransactionValue().isEmpty()) {
        transactionValue_ = other.transactionValue_;
        onChanged();
      }
      if (other.getBlockTimestamp() != 0L) {
        setBlockTimestamp(other.getBlockTimestamp());
      }
      this.mergeUnknownFields(other.getUnknownFields());
      onChanged();
      return this;
    }

    @java.lang.Override
    public final boolean isInitialized() {
      return true;
    }

    @java.lang.Override
    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              transactionHash_ = input.readStringRequireUtf8();

              break;
            } // case 10
            case 16: {
              transactionIndex_ = input.readInt32();

              break;
            } // case 16
            case 24: {
              blockNumber_ = input.readInt64();

              break;
            } // case 24
            case 34: {
              fromAddress_ = input.readStringRequireUtf8();

              break;
            } // case 34
            case 42: {
              toAddress_ = input.readStringRequireUtf8();

              break;
            } // case 42
            case 48: {
              nonce_ = input.readInt64();

              break;
            } // case 48
            case 58: {
              transactionValue_ = input.readStringRequireUtf8();

              break;
            } // case 58
            case 64: {
              blockTimestamp_ = input.readInt64();

              break;
            } // case 64
            default: {
              if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                done = true; // was an endgroup tag
              }
              break;
            } // default:
          } // switch (tag)
        } // while (!done)
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.unwrapIOException();
      } finally {
        onChanged();
      } // finally
      return this;
    }

    private java.lang.Object transactionHash_ = "";
    /**
     * <code>string transactionHash = 1;</code>
     * @return The transactionHash.
     */
    public java.lang.String getTransactionHash() {
      java.lang.Object ref = transactionHash_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        transactionHash_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string transactionHash = 1;</code>
     * @return The bytes for transactionHash.
     */
    public com.google.protobuf.ByteString
        getTransactionHashBytes() {
      java.lang.Object ref = transactionHash_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        transactionHash_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string transactionHash = 1;</code>
     * @param value The transactionHash to set.
     * @return This builder for chaining.
     */
    public Builder setTransactionHash(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      transactionHash_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string transactionHash = 1;</code>
     * @return This builder for chaining.
     */
    public Builder clearTransactionHash() {
      
      transactionHash_ = getDefaultInstance().getTransactionHash();
      onChanged();
      return this;
    }
    /**
     * <code>string transactionHash = 1;</code>
     * @param value The bytes for transactionHash to set.
     * @return This builder for chaining.
     */
    public Builder setTransactionHashBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      transactionHash_ = value;
      onChanged();
      return this;
    }

    private int transactionIndex_ ;
    /**
     * <code>int32 transactionIndex = 2;</code>
     * @return The transactionIndex.
     */
    @java.lang.Override
    public int getTransactionIndex() {
      return transactionIndex_;
    }
    /**
     * <code>int32 transactionIndex = 2;</code>
     * @param value The transactionIndex to set.
     * @return This builder for chaining.
     */
    public Builder setTransactionIndex(int value) {
      
      transactionIndex_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 transactionIndex = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearTransactionIndex() {
      
      transactionIndex_ = 0;
      onChanged();
      return this;
    }

    private long blockNumber_ ;
    /**
     * <code>int64 blockNumber = 3;</code>
     * @return The blockNumber.
     */
    @java.lang.Override
    public long getBlockNumber() {
      return blockNumber_;
    }
    /**
     * <code>int64 blockNumber = 3;</code>
     * @param value The blockNumber to set.
     * @return This builder for chaining.
     */
    public Builder setBlockNumber(long value) {
      
      blockNumber_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 blockNumber = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearBlockNumber() {
      
      blockNumber_ = 0L;
      onChanged();
      return this;
    }

    private java.lang.Object fromAddress_ = "";
    /**
     * <code>string fromAddress = 4;</code>
     * @return The fromAddress.
     */
    public java.lang.String getFromAddress() {
      java.lang.Object ref = fromAddress_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        fromAddress_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string fromAddress = 4;</code>
     * @return The bytes for fromAddress.
     */
    public com.google.protobuf.ByteString
        getFromAddressBytes() {
      java.lang.Object ref = fromAddress_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        fromAddress_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string fromAddress = 4;</code>
     * @param value The fromAddress to set.
     * @return This builder for chaining.
     */
    public Builder setFromAddress(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      fromAddress_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string fromAddress = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearFromAddress() {
      
      fromAddress_ = getDefaultInstance().getFromAddress();
      onChanged();
      return this;
    }
    /**
     * <code>string fromAddress = 4;</code>
     * @param value The bytes for fromAddress to set.
     * @return This builder for chaining.
     */
    public Builder setFromAddressBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      fromAddress_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object toAddress_ = "";
    /**
     * <code>string toAddress = 5;</code>
     * @return The toAddress.
     */
    public java.lang.String getToAddress() {
      java.lang.Object ref = toAddress_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        toAddress_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string toAddress = 5;</code>
     * @return The bytes for toAddress.
     */
    public com.google.protobuf.ByteString
        getToAddressBytes() {
      java.lang.Object ref = toAddress_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        toAddress_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string toAddress = 5;</code>
     * @param value The toAddress to set.
     * @return This builder for chaining.
     */
    public Builder setToAddress(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      toAddress_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string toAddress = 5;</code>
     * @return This builder for chaining.
     */
    public Builder clearToAddress() {
      
      toAddress_ = getDefaultInstance().getToAddress();
      onChanged();
      return this;
    }
    /**
     * <code>string toAddress = 5;</code>
     * @param value The bytes for toAddress to set.
     * @return This builder for chaining.
     */
    public Builder setToAddressBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      toAddress_ = value;
      onChanged();
      return this;
    }

    private long nonce_ ;
    /**
     * <code>int64 nonce = 6;</code>
     * @return The nonce.
     */
    @java.lang.Override
    public long getNonce() {
      return nonce_;
    }
    /**
     * <code>int64 nonce = 6;</code>
     * @param value The nonce to set.
     * @return This builder for chaining.
     */
    public Builder setNonce(long value) {
      
      nonce_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 nonce = 6;</code>
     * @return This builder for chaining.
     */
    public Builder clearNonce() {
      
      nonce_ = 0L;
      onChanged();
      return this;
    }

    private java.lang.Object transactionValue_ = "";
    /**
     * <code>string transactionValue = 7;</code>
     * @return The transactionValue.
     */
    public java.lang.String getTransactionValue() {
      java.lang.Object ref = transactionValue_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        transactionValue_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string transactionValue = 7;</code>
     * @return The bytes for transactionValue.
     */
    public com.google.protobuf.ByteString
        getTransactionValueBytes() {
      java.lang.Object ref = transactionValue_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        transactionValue_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string transactionValue = 7;</code>
     * @param value The transactionValue to set.
     * @return This builder for chaining.
     */
    public Builder setTransactionValue(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      transactionValue_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string transactionValue = 7;</code>
     * @return This builder for chaining.
     */
    public Builder clearTransactionValue() {
      
      transactionValue_ = getDefaultInstance().getTransactionValue();
      onChanged();
      return this;
    }
    /**
     * <code>string transactionValue = 7;</code>
     * @param value The bytes for transactionValue to set.
     * @return This builder for chaining.
     */
    public Builder setTransactionValueBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      transactionValue_ = value;
      onChanged();
      return this;
    }

    private long blockTimestamp_ ;
    /**
     * <code>int64 blockTimestamp = 8;</code>
     * @return The blockTimestamp.
     */
    @java.lang.Override
    public long getBlockTimestamp() {
      return blockTimestamp_;
    }
    /**
     * <code>int64 blockTimestamp = 8;</code>
     * @param value The blockTimestamp to set.
     * @return This builder for chaining.
     */
    public Builder setBlockTimestamp(long value) {
      
      blockTimestamp_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 blockTimestamp = 8;</code>
     * @return This builder for chaining.
     */
    public Builder clearBlockTimestamp() {
      
      blockTimestamp_ = 0L;
      onChanged();
      return this;
    }
    @java.lang.Override
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.setUnknownFields(unknownFields);
    }

    @java.lang.Override
    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return super.mergeUnknownFields(unknownFields);
    }


    // @@protoc_insertion_point(builder_scope:erigon.Transaction)
  }

  // @@protoc_insertion_point(class_scope:erigon.Transaction)
  private static final io.eigenphi.erigon.protos.Transaction DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.eigenphi.erigon.protos.Transaction();
  }

  public static io.eigenphi.erigon.protos.Transaction getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Transaction>
      PARSER = new com.google.protobuf.AbstractParser<Transaction>() {
    @java.lang.Override
    public Transaction parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      Builder builder = newBuilder();
      try {
        builder.mergeFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(builder.buildPartial());
      } catch (com.google.protobuf.UninitializedMessageException e) {
        throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(e)
            .setUnfinishedMessage(builder.buildPartial());
      }
      return builder.buildPartial();
    }
  };

  public static com.google.protobuf.Parser<Transaction> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Transaction> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.eigenphi.erigon.protos.Transaction getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

