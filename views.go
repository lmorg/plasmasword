package main

const (
	// default
	viewAll = `CREATE VIEW IF NOT EXISTS default_view AS
		SELECT datetime, status, size, proc, uri
		FROM access
		ORDER BY datetime DESC;`

	// latest:
	viewLatestNon200 = `CREATE VIEW latest_non_200 AS
		SELECT datetime, status, uri
		FROM access
		WHERE status != 200
		ORDER BY datetime DESC;`

	viewLatestProc = `CREATE VIEW latest_proc AS
		SELECT datetime, status, proc, uri
		FROM access
		WHERE status = 200
		ORDER BY datetime DESC;`

	viewLatest304 = `CREATE VIEW latest_304 AS
		SELECT status, datetime, uri, qs
		FROM access
		WHERE status = 304
		ORDER BY datetime DESC;`

	// counts:
	viewCountStatus = `CREATE VIEW count_status AS
		SELECT count(*) "#", status
		FROM access
		GROUP BY status
		ORDER BY 1 DESC;`

	viewCount304 = `CREATE VIEW count_304 AS
		SELECT count(status) "#", uri
		FROM access
		WHERE status = 304
		GROUP BY uri
		ORDER BY 1 DESC;`

	viewCountSize = `CREATE VIEW count_size AS
		SELECT size, count(uri) "#", uri, qs
		FROM access
		GROUP BY size
		ORDER BY 1 DESC;`

	// misc
	viewListViews = `CREATE VIEW list_views AS
		SELECT sql FROM sqlite_master WHERE type = 'view';`
)
