package raft.models;

import raft.rpc.RequestVoteRPC;
import raft.util.Role;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.Random;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicInteger;

public class RaftNode {
    private final int nodeId;
    private final List<Integer> peers;
    private int currentTerm;
    private Integer votedFor;
    private final List<LogEntry> log = new ArrayList<>();
    private Role role = Role.FOLLOWER;
    private int commitIndex = -1;
    private int lastApplied = -1;
    private Map<Integer, Integer> nextIndex;
    private Map<Integer, Integer> matchIndex;
    private final ScheduledExecutorService scheduler = Executors.newScheduledThreadPool(2);
    private ScheduledFuture<?> electionTimeout;
    private ScheduledFuture<?> heartbeat;

    public RaftNode(int nodeId, List<Integer> peers) {
        this.nodeId = nodeId;
        this.peers = peers;
    }
}
