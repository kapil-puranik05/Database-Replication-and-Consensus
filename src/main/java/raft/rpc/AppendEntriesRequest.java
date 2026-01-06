package raft.rpc;

import raft.models.LogEntry;

import java.util.List;

public class AppendEntriesRequest {
    public int term;
    public int leaderId;
    public int prevLogIndex;
    public int prevLogTerm;
    public List<LogEntry> entries;
    public int leaderCommit;
}
