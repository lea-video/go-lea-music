CREATE TABLE IF NOT EXISTS info (
    key TEXT,
    val TEXT
);
INSERT INTO info (key, val) VALUES ("version", "1");

-------------
-- Artists --
-------------

CREATE TABLE IF NOT EXISTS artists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    is_group BOOLEAN
);

CREATE TABLE IF NOT EXISTS map_artist_group_members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    artist_group INTEGER,
    member INTEGER,
    FOREIGN KEY(artist_group) REFERENCES artists(id),
    FOREIGN KEY(member) REFERENCES artists(id)
);

CREATE VIEW view_artist_groups AS
    SELECT a.id, a.name, is_group, GROUP_CONCAT(m.member) AS members
    FROM artists AS a
        LEFT JOIN map_artist_group_members AS m ON a.id = m.artist_group
    WHERE a.is_group = 1
    GROUP BY a.id, a.name;

-----------
-- Media --
-----------

CREATE TABLE IF NOT EXISTS media (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source TEXT
);

CREATE TABLE IF NOT EXISTS mediatrack (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parentMedia INTEGER,
    has_audio BOOLEAN,
    has_video BOOLEAN,
    has_picture BOOLEAN,
    FOREIGN KEY(parentMedia) REFERENCES media(id)
);

CREATE VIEW view_media_tracks AS
    SELECT m.id, m.source, GROUP_CONCAT(mt.id) AS tracks
    FROM media AS m
        LEFT JOIN mediatrack AS mt ON m.id = mt.parentMedia
    GROUP BY m.id, m.source;

CREATE TABLE IF NOT EXISTS tmpfile (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parentMedia INTEGER,
    accessToken TEXT,
    location TEXT,
    maxage INTEGER,
    FOREIGN KEY(parentMedia) REFERENCES media(id)
);
