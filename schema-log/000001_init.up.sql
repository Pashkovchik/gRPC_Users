CREATE TABLE logs
(
    date DATE,
    size Int32,
    message String
) ENGINE = MergeTree(date, message, 8192);