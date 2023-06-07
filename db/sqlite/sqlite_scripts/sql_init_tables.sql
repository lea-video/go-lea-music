CREATE TABLE IF NOT EXISTS info (
    key TEXT NOT NULL,
    val TEXT NOT NULL
);
INSERT INTO info (key, val) VALUES ("version", "1");

-------------
-- Artists --
-------------

CREATE TABLE IF NOT EXISTS artists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    is_group BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS map_artist_group_members (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    artist_group INTEGER NOT NULL,
    member INTEGER NOT NULL,
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
    source TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS media_track (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parentMedia INTEGER NOT NULL,
    has_audio BOOLEAN NOT NULL,
    has_video BOOLEAN NOT NULL,
    has_picture BOOLEAN NOT NULL,
    FOREIGN KEY(parentMedia) REFERENCES media(id)
);

CREATE VIEW view_media_tracks AS
    SELECT m.id, m.source, GROUP_CONCAT(mt.id) AS tracks
    FROM media AS m
        LEFT JOIN media_track AS mt ON m.id = mt.parentMedia
    GROUP BY m.id, m.source;

CREATE TABLE IF NOT EXISTS tmp_file (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    parentMedia INTEGER NOT NULL,
    accessToken TEXT NOT NULL,
    location TEXT NOT NULL,
    maxAge INTEGER NOT NULL,
    FOREIGN KEY(parentMedia) REFERENCES media(id)
);

---------------
-- Playlists --
---------------

CREATE TABLE IF NOT EXISTS playlists (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS playlist_item (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    playlist INTEGER NOT NULL,
    position REAL NOT NULL,
    element INTEGER NOT NULL,
    FOREIGN KEY(playlist) REFERENCES playlists(id)
);

CREATE VIEW view_playlists_with_items AS
    SELECT pl.id, pl.name, GROUP_CONCAT(pli.element) AS items
    FROM playlists AS pl
        LEFT JOIN playlist_item AS pli ON pl.id = pli.playlist
    GROUP BY pl.id, pl.name;
