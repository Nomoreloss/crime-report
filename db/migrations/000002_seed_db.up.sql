
INSERT INTO role (id, name, description, access_level, created_at,updated_at) VALUES
('a9c839e9-d3d1-4177-8014-d36eba8d6a00', 'Admin', 'Admin role', 1, '2022-11-10T23:00:00Z','2022-11-10T23:00:00Z'),
('edd35257-6ced-485b-8681-bd86ddf46413', 'Staff', 'Staff role', 2, '2022-11-10T23:00:00Z','2022-11-10T23:00:00Z'),
('d7d4146e-4dbc-4390-9463-55e25297c437', 'User', 'User role', 3,  '2022-11-10T23:00:00Z','2022-11-10T23:00:00Z');


INSERT INTO users (id, username,email,password,first_name,last_name,user_type,mobile,status,role,active,created_at,updated_at) VALUES
('d7d4146e-4dbc-4390-9463-55e25297c437','admin','admin@admin.com','%s', 'John','Doe','Admin','08090909898','ACTIVE', 'a9c839e9-d3d1-4177-8014-d36eba8d6a00',true, '2022-11-10T23:00:00Z','2022-11-10T23:00:00Z'),
('bdd2006a-cfa2-4972-9f56-051c58e22fb5','staff','staf@staff.com', '%s','Stella','Amos','Admin','08090909898','ACTIVE','edd35257-6ced-485b-8681-bd86ddf46413', true,'2022-11-10T23:00:00Z','2022-11-10T23:00:00Z');
