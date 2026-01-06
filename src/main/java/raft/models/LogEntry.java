package raft.models;

public class LogEntry {
    public final int term;
    public final String command;

    public LogEntry(int term, String command) {
        this.term = term;
        this.command = command;
    }
}
