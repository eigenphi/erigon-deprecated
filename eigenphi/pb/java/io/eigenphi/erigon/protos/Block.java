// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: Block.proto

package io.eigenphi.erigon.protos;

/**
 * Protobuf type {@code erigon.Block}
 */
public final class Block extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:erigon.Block)
    BlockOrBuilder {
private static final long serialVersionUID = 0L;
  // Use Block.newBuilder() to construct.
  private Block(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private Block() {
    blockHash_ = "";
    parentHash_ = "";
    miner_ = "";
  }

  @java.lang.Override
  @SuppressWarnings({"unused"})
  protected java.lang.Object newInstance(
      UnusedPrivateParameter unused) {
    return new Block();
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return this.unknownFields;
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Block_descriptor;
  }

  @java.lang.Override
  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Block_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            io.eigenphi.erigon.protos.Block.class, io.eigenphi.erigon.protos.Block.Builder.class);
  }

  public static final int BLOCKNUMBER_FIELD_NUMBER = 2;
  private long blockNumber_;
  /**
   * <code>int64 blockNumber = 2;</code>
   * @return The blockNumber.
   */
  @java.lang.Override
  public long getBlockNumber() {
    return blockNumber_;
  }

  public static final int BLOCKHASH_FIELD_NUMBER = 3;
  private volatile java.lang.Object blockHash_;
  /**
   * <code>string blockHash = 3;</code>
   * @return The blockHash.
   */
  @java.lang.Override
  public java.lang.String getBlockHash() {
    java.lang.Object ref = blockHash_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      blockHash_ = s;
      return s;
    }
  }
  /**
   * <code>string blockHash = 3;</code>
   * @return The bytes for blockHash.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getBlockHashBytes() {
    java.lang.Object ref = blockHash_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      blockHash_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int PARENTHASH_FIELD_NUMBER = 4;
  private volatile java.lang.Object parentHash_;
  /**
   * <code>string parentHash = 4;</code>
   * @return The parentHash.
   */
  @java.lang.Override
  public java.lang.String getParentHash() {
    java.lang.Object ref = parentHash_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      parentHash_ = s;
      return s;
    }
  }
  /**
   * <code>string parentHash = 4;</code>
   * @return The bytes for parentHash.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getParentHashBytes() {
    java.lang.Object ref = parentHash_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      parentHash_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int MINER_FIELD_NUMBER = 6;
  private volatile java.lang.Object miner_;
  /**
   * <code>string miner = 6;</code>
   * @return The miner.
   */
  @java.lang.Override
  public java.lang.String getMiner() {
    java.lang.Object ref = miner_;
    if (ref instanceof java.lang.String) {
      return (java.lang.String) ref;
    } else {
      com.google.protobuf.ByteString bs = 
          (com.google.protobuf.ByteString) ref;
      java.lang.String s = bs.toStringUtf8();
      miner_ = s;
      return s;
    }
  }
  /**
   * <code>string miner = 6;</code>
   * @return The bytes for miner.
   */
  @java.lang.Override
  public com.google.protobuf.ByteString
      getMinerBytes() {
    java.lang.Object ref = miner_;
    if (ref instanceof java.lang.String) {
      com.google.protobuf.ByteString b = 
          com.google.protobuf.ByteString.copyFromUtf8(
              (java.lang.String) ref);
      miner_ = b;
      return b;
    } else {
      return (com.google.protobuf.ByteString) ref;
    }
  }

  public static final int BLOCKSIZE_FIELD_NUMBER = 7;
  private int blockSize_;
  /**
   * <code>int32 blockSize = 7;</code>
   * @return The blockSize.
   */
  @java.lang.Override
  public int getBlockSize() {
    return blockSize_;
  }

  public static final int GASLIMIT_FIELD_NUMBER = 8;
  private long gasLimit_;
  /**
   * <code>int64 gasLimit = 8;</code>
   * @return The gasLimit.
   */
  @java.lang.Override
  public long getGasLimit() {
    return gasLimit_;
  }

  public static final int GASUSED_FIELD_NUMBER = 9;
  private long gasUsed_;
  /**
   * <code>int64 gasUsed = 9;</code>
   * @return The gasUsed.
   */
  @java.lang.Override
  public long getGasUsed() {
    return gasUsed_;
  }

  public static final int BLOCKTIMESTAMP_FIELD_NUMBER = 10;
  private long blockTimestamp_;
  /**
   * <code>int64 blockTimestamp = 10;</code>
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
    if (blockNumber_ != 0L) {
      output.writeInt64(2, blockNumber_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(blockHash_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 3, blockHash_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(parentHash_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 4, parentHash_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(miner_)) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 6, miner_);
    }
    if (blockSize_ != 0) {
      output.writeInt32(7, blockSize_);
    }
    if (gasLimit_ != 0L) {
      output.writeInt64(8, gasLimit_);
    }
    if (gasUsed_ != 0L) {
      output.writeInt64(9, gasUsed_);
    }
    if (blockTimestamp_ != 0L) {
      output.writeInt64(10, blockTimestamp_);
    }
    getUnknownFields().writeTo(output);
  }

  @java.lang.Override
  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    if (blockNumber_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(2, blockNumber_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(blockHash_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, blockHash_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(parentHash_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(4, parentHash_);
    }
    if (!com.google.protobuf.GeneratedMessageV3.isStringEmpty(miner_)) {
      size += com.google.protobuf.GeneratedMessageV3.computeStringSize(6, miner_);
    }
    if (blockSize_ != 0) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt32Size(7, blockSize_);
    }
    if (gasLimit_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(8, gasLimit_);
    }
    if (gasUsed_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(9, gasUsed_);
    }
    if (blockTimestamp_ != 0L) {
      size += com.google.protobuf.CodedOutputStream
        .computeInt64Size(10, blockTimestamp_);
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
    if (!(obj instanceof io.eigenphi.erigon.protos.Block)) {
      return super.equals(obj);
    }
    io.eigenphi.erigon.protos.Block other = (io.eigenphi.erigon.protos.Block) obj;

    if (getBlockNumber()
        != other.getBlockNumber()) return false;
    if (!getBlockHash()
        .equals(other.getBlockHash())) return false;
    if (!getParentHash()
        .equals(other.getParentHash())) return false;
    if (!getMiner()
        .equals(other.getMiner())) return false;
    if (getBlockSize()
        != other.getBlockSize()) return false;
    if (getGasLimit()
        != other.getGasLimit()) return false;
    if (getGasUsed()
        != other.getGasUsed()) return false;
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
    hash = (37 * hash) + BLOCKNUMBER_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getBlockNumber());
    hash = (37 * hash) + BLOCKHASH_FIELD_NUMBER;
    hash = (53 * hash) + getBlockHash().hashCode();
    hash = (37 * hash) + PARENTHASH_FIELD_NUMBER;
    hash = (53 * hash) + getParentHash().hashCode();
    hash = (37 * hash) + MINER_FIELD_NUMBER;
    hash = (53 * hash) + getMiner().hashCode();
    hash = (37 * hash) + BLOCKSIZE_FIELD_NUMBER;
    hash = (53 * hash) + getBlockSize();
    hash = (37 * hash) + GASLIMIT_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getGasLimit());
    hash = (37 * hash) + GASUSED_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getGasUsed());
    hash = (37 * hash) + BLOCKTIMESTAMP_FIELD_NUMBER;
    hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
        getBlockTimestamp());
    hash = (29 * hash) + getUnknownFields().hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static io.eigenphi.erigon.protos.Block parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Block parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static io.eigenphi.erigon.protos.Block parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static io.eigenphi.erigon.protos.Block parseFrom(
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
  public static Builder newBuilder(io.eigenphi.erigon.protos.Block prototype) {
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
   * Protobuf type {@code erigon.Block}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:erigon.Block)
      io.eigenphi.erigon.protos.BlockOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Block_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Block_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.eigenphi.erigon.protos.Block.class, io.eigenphi.erigon.protos.Block.Builder.class);
    }

    // Construct using io.eigenphi.erigon.protos.Block.newBuilder()
    private Builder() {

    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);

    }
    @java.lang.Override
    public Builder clear() {
      super.clear();
      blockNumber_ = 0L;

      blockHash_ = "";

      parentHash_ = "";

      miner_ = "";

      blockSize_ = 0;

      gasLimit_ = 0L;

      gasUsed_ = 0L;

      blockTimestamp_ = 0L;

      return this;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return io.eigenphi.erigon.protos.ErigonProtos.internal_static_erigon_Block_descriptor;
    }

    @java.lang.Override
    public io.eigenphi.erigon.protos.Block getDefaultInstanceForType() {
      return io.eigenphi.erigon.protos.Block.getDefaultInstance();
    }

    @java.lang.Override
    public io.eigenphi.erigon.protos.Block build() {
      io.eigenphi.erigon.protos.Block result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    @java.lang.Override
    public io.eigenphi.erigon.protos.Block buildPartial() {
      io.eigenphi.erigon.protos.Block result = new io.eigenphi.erigon.protos.Block(this);
      result.blockNumber_ = blockNumber_;
      result.blockHash_ = blockHash_;
      result.parentHash_ = parentHash_;
      result.miner_ = miner_;
      result.blockSize_ = blockSize_;
      result.gasLimit_ = gasLimit_;
      result.gasUsed_ = gasUsed_;
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
      if (other instanceof io.eigenphi.erigon.protos.Block) {
        return mergeFrom((io.eigenphi.erigon.protos.Block)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(io.eigenphi.erigon.protos.Block other) {
      if (other == io.eigenphi.erigon.protos.Block.getDefaultInstance()) return this;
      if (other.getBlockNumber() != 0L) {
        setBlockNumber(other.getBlockNumber());
      }
      if (!other.getBlockHash().isEmpty()) {
        blockHash_ = other.blockHash_;
        onChanged();
      }
      if (!other.getParentHash().isEmpty()) {
        parentHash_ = other.parentHash_;
        onChanged();
      }
      if (!other.getMiner().isEmpty()) {
        miner_ = other.miner_;
        onChanged();
      }
      if (other.getBlockSize() != 0) {
        setBlockSize(other.getBlockSize());
      }
      if (other.getGasLimit() != 0L) {
        setGasLimit(other.getGasLimit());
      }
      if (other.getGasUsed() != 0L) {
        setGasUsed(other.getGasUsed());
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
            case 16: {
              blockNumber_ = input.readInt64();

              break;
            } // case 16
            case 26: {
              blockHash_ = input.readStringRequireUtf8();

              break;
            } // case 26
            case 34: {
              parentHash_ = input.readStringRequireUtf8();

              break;
            } // case 34
            case 50: {
              miner_ = input.readStringRequireUtf8();

              break;
            } // case 50
            case 56: {
              blockSize_ = input.readInt32();

              break;
            } // case 56
            case 64: {
              gasLimit_ = input.readInt64();

              break;
            } // case 64
            case 72: {
              gasUsed_ = input.readInt64();

              break;
            } // case 72
            case 80: {
              blockTimestamp_ = input.readInt64();

              break;
            } // case 80
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

    private long blockNumber_ ;
    /**
     * <code>int64 blockNumber = 2;</code>
     * @return The blockNumber.
     */
    @java.lang.Override
    public long getBlockNumber() {
      return blockNumber_;
    }
    /**
     * <code>int64 blockNumber = 2;</code>
     * @param value The blockNumber to set.
     * @return This builder for chaining.
     */
    public Builder setBlockNumber(long value) {
      
      blockNumber_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 blockNumber = 2;</code>
     * @return This builder for chaining.
     */
    public Builder clearBlockNumber() {
      
      blockNumber_ = 0L;
      onChanged();
      return this;
    }

    private java.lang.Object blockHash_ = "";
    /**
     * <code>string blockHash = 3;</code>
     * @return The blockHash.
     */
    public java.lang.String getBlockHash() {
      java.lang.Object ref = blockHash_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        blockHash_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string blockHash = 3;</code>
     * @return The bytes for blockHash.
     */
    public com.google.protobuf.ByteString
        getBlockHashBytes() {
      java.lang.Object ref = blockHash_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        blockHash_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string blockHash = 3;</code>
     * @param value The blockHash to set.
     * @return This builder for chaining.
     */
    public Builder setBlockHash(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      blockHash_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string blockHash = 3;</code>
     * @return This builder for chaining.
     */
    public Builder clearBlockHash() {
      
      blockHash_ = getDefaultInstance().getBlockHash();
      onChanged();
      return this;
    }
    /**
     * <code>string blockHash = 3;</code>
     * @param value The bytes for blockHash to set.
     * @return This builder for chaining.
     */
    public Builder setBlockHashBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      blockHash_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object parentHash_ = "";
    /**
     * <code>string parentHash = 4;</code>
     * @return The parentHash.
     */
    public java.lang.String getParentHash() {
      java.lang.Object ref = parentHash_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        parentHash_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string parentHash = 4;</code>
     * @return The bytes for parentHash.
     */
    public com.google.protobuf.ByteString
        getParentHashBytes() {
      java.lang.Object ref = parentHash_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        parentHash_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string parentHash = 4;</code>
     * @param value The parentHash to set.
     * @return This builder for chaining.
     */
    public Builder setParentHash(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      parentHash_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string parentHash = 4;</code>
     * @return This builder for chaining.
     */
    public Builder clearParentHash() {
      
      parentHash_ = getDefaultInstance().getParentHash();
      onChanged();
      return this;
    }
    /**
     * <code>string parentHash = 4;</code>
     * @param value The bytes for parentHash to set.
     * @return This builder for chaining.
     */
    public Builder setParentHashBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      parentHash_ = value;
      onChanged();
      return this;
    }

    private java.lang.Object miner_ = "";
    /**
     * <code>string miner = 6;</code>
     * @return The miner.
     */
    public java.lang.String getMiner() {
      java.lang.Object ref = miner_;
      if (!(ref instanceof java.lang.String)) {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        miner_ = s;
        return s;
      } else {
        return (java.lang.String) ref;
      }
    }
    /**
     * <code>string miner = 6;</code>
     * @return The bytes for miner.
     */
    public com.google.protobuf.ByteString
        getMinerBytes() {
      java.lang.Object ref = miner_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        miner_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }
    /**
     * <code>string miner = 6;</code>
     * @param value The miner to set.
     * @return This builder for chaining.
     */
    public Builder setMiner(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  
      miner_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>string miner = 6;</code>
     * @return This builder for chaining.
     */
    public Builder clearMiner() {
      
      miner_ = getDefaultInstance().getMiner();
      onChanged();
      return this;
    }
    /**
     * <code>string miner = 6;</code>
     * @param value The bytes for miner to set.
     * @return This builder for chaining.
     */
    public Builder setMinerBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      
      miner_ = value;
      onChanged();
      return this;
    }

    private int blockSize_ ;
    /**
     * <code>int32 blockSize = 7;</code>
     * @return The blockSize.
     */
    @java.lang.Override
    public int getBlockSize() {
      return blockSize_;
    }
    /**
     * <code>int32 blockSize = 7;</code>
     * @param value The blockSize to set.
     * @return This builder for chaining.
     */
    public Builder setBlockSize(int value) {
      
      blockSize_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int32 blockSize = 7;</code>
     * @return This builder for chaining.
     */
    public Builder clearBlockSize() {
      
      blockSize_ = 0;
      onChanged();
      return this;
    }

    private long gasLimit_ ;
    /**
     * <code>int64 gasLimit = 8;</code>
     * @return The gasLimit.
     */
    @java.lang.Override
    public long getGasLimit() {
      return gasLimit_;
    }
    /**
     * <code>int64 gasLimit = 8;</code>
     * @param value The gasLimit to set.
     * @return This builder for chaining.
     */
    public Builder setGasLimit(long value) {
      
      gasLimit_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 gasLimit = 8;</code>
     * @return This builder for chaining.
     */
    public Builder clearGasLimit() {
      
      gasLimit_ = 0L;
      onChanged();
      return this;
    }

    private long gasUsed_ ;
    /**
     * <code>int64 gasUsed = 9;</code>
     * @return The gasUsed.
     */
    @java.lang.Override
    public long getGasUsed() {
      return gasUsed_;
    }
    /**
     * <code>int64 gasUsed = 9;</code>
     * @param value The gasUsed to set.
     * @return This builder for chaining.
     */
    public Builder setGasUsed(long value) {
      
      gasUsed_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 gasUsed = 9;</code>
     * @return This builder for chaining.
     */
    public Builder clearGasUsed() {
      
      gasUsed_ = 0L;
      onChanged();
      return this;
    }

    private long blockTimestamp_ ;
    /**
     * <code>int64 blockTimestamp = 10;</code>
     * @return The blockTimestamp.
     */
    @java.lang.Override
    public long getBlockTimestamp() {
      return blockTimestamp_;
    }
    /**
     * <code>int64 blockTimestamp = 10;</code>
     * @param value The blockTimestamp to set.
     * @return This builder for chaining.
     */
    public Builder setBlockTimestamp(long value) {
      
      blockTimestamp_ = value;
      onChanged();
      return this;
    }
    /**
     * <code>int64 blockTimestamp = 10;</code>
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


    // @@protoc_insertion_point(builder_scope:erigon.Block)
  }

  // @@protoc_insertion_point(class_scope:erigon.Block)
  private static final io.eigenphi.erigon.protos.Block DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new io.eigenphi.erigon.protos.Block();
  }

  public static io.eigenphi.erigon.protos.Block getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<Block>
      PARSER = new com.google.protobuf.AbstractParser<Block>() {
    @java.lang.Override
    public Block parsePartialFrom(
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

  public static com.google.protobuf.Parser<Block> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<Block> getParserForType() {
    return PARSER;
  }

  @java.lang.Override
  public io.eigenphi.erigon.protos.Block getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

