insert into public.segments (id, name)
values  (1, 'Valencia St'),
        (2, 'Castro St'),
        (3, 'Montgomery St'),
        (4, 'Market St');


insert into public.regulations (id, name, segment_id)
values  (1, 'Street Cleaning', 1),
        (2, 'Tow Away', 1),
        (3, 'Time Limit', 1),
        (4, 'Meters', 2),
        (5, 'Time Limit', 2),
        (6, 'Tow Away', 2),
        (7, 'Meters', 3),
        (8, 'Tow Away', 3),
        (9, 'Street Cleaning', 4),
        (10, 'Time Limit', 4);


insert into public.regulated_slots (id, day_of_week, week_of_month, day_of_month, start_time, end_time, allowed_threshold, cost, is_daily, regulation_id)
values  (3, null, null, null, '10:00:00', '14:00:00', null, 0, true, 2),
        (6, 3, null, null, '09:00:00', '18:00:00', 2, 0, false, 3),
        (2, 5, 4, null, '09:00:00', '11:00:00', null, 0, false, 1),
        (1, 5, 2, null, '09:00:00', '11:00:00', null, 0, false, 1),
        (7, 4, null, null, '09:00:00', '18:00:00', 2, 0, false, 3),
        (4, 1, null, null, '09:00:00', '18:00:00', 2, 0, false, 3),
        (8, 5, null, null, '09:00:00', '18:00:00', 2, 0, false, 3),
        (5, 2, null, null, '09:00:00', '18:00:00', 2, 0, false, 3),
        (9, null, null, null, '08:00:00', '18:00:00', 2, 5, true, 4),
        (10, 4, null, null, '09:00:00', '18:00:00', 3, 0, false, 5),
        (11, 5, null, null, '09:00:00', '18:00:00', 3, 0, false, 5),
        (12, 1, null, null, '14:00:00', '16:00:00', null, 0, false, 6),
        (13, 2, null, null, '14:00:00', '16:00:00', null, 0, false, 6),
        (14, null, null, null, '08:00:00', '18:00:00', 2, 5, true, 7),
        (15, 4, null, null, '14:00:00', '16:00:00', null, 0, false, 8),
        (16, 5, 2, null, '09:00:00', '11:00:00', null, 0, false, 9),
        (17, 5, 4, null, '09:00:00', '11:00:00', null, 0, false, 9),
        (18, null, null, null, '09:00:00', '18:00:00', null, 0, true, 10);