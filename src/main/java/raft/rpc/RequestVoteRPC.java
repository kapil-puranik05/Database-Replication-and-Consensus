package raft.rpc;

public class RequestVoteRPC {
    public int term;
    public int candidateId;
    public int lastLogIndex;
    public int lastLogTerm;
}
