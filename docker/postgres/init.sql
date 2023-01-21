SELECT 'CREATE DATABASE dartsly_match_service'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'dartsly_match_service')\gexec