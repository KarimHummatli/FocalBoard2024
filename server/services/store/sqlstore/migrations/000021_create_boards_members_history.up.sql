{{if doesTableExist .schemaName "board_members_history" }}

SELECT 1;

{{else}}

CREATE TABLE IF NOT EXISTS {{.prefix}}board_members_history (
    board_id VARCHAR(36) NOT NULL,
    user_id VARCHAR(36) NOT NULL,
    action VARCHAR(10),
    {{if .postgres}}insert_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),{{end}}
	{{if .sqlite}}insert_at DATETIME NOT NULL DEFAULT(STRFTIME('%Y-%m-%d %H:%M:%f', 'NOW')),{{end}}
	{{if .mysql}}insert_at DATETIME(6) NOT NULL DEFAULT NOW(6),{{end}}
    PRIMARY KEY (board_id, user_id, insert_at)
) {{if .mysql}}DEFAULT CHARACTER SET utf8mb4{{end}};

INSERT INTO {{.prefix}}board_members_history (board_id, user_id, action) SELECT board_id, user_id, 'created' from {{.prefix}}board_members;

{{end}}

{{- /* createIndexIfNeeded schemaName tableName columns */ -}}
{{ createIndexIfNeeded .schemaName "board_members_history" "user_id" }}
{{ createIndexIfNeeded .schemaName "board_members_history" "board_id, user_id" }}
