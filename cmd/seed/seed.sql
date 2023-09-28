insert into currencies (id,code) values ('a111b4d8-23fd-4654-b73e-47b87c8e03aa','KZT'),
                                        ('b506a1c9-7661-4f9e-9e1b-dbd6c978a4f5','KGS');
insert into countries (id, code) values ('ebd63655-0408-4976-8bfb-2d70d183caf1', 'KAZ'),
                                        ('d3eee6d5-0e70-4c02-9837-d6ecdb19b7eb', 'KGZ');

insert into public.wallets (id,phone, currency_id, country_id, balance, hold, status, identification,created_at)
values  ('01875aa2-b397-a341-e8e8-9541ff42518e', null, 'a111b4d8-23fd-4654-b73e-47b87c8e03aa', 'ebd63655-0408-4976-8bfb-2d70d183caf1', 0, 0, 'active','basic','2023-04-07 07:33:44.974579'),
        ('12975b22-ca2a-7b87-fc49-9d20e1f38d83', 123, 'a111b4d8-23fd-4654-b73e-47b87c8e03aa', 'ebd63655-0408-4976-8bfb-2d70d183caf1', 0, 0, 'active', 'full','2023-04-07 09:53:39.362652'),
        ('234769c7-0718-5b1d-2612-277b1980364c', null, 'a111b4d8-23fd-4654-b73e-47b87c8e03aa', 'ebd63655-0408-4976-8bfb-2d70d183caf1', 0, 0, 'active', 'none','2023-04-10 06:07:43.895930');

insert into public.consumers (id, code, slug, secret, white_list_methods, created_at) values ('9aff1dd3-f474-4a12-ab5d-8c28110eade8', 'dog1', 'myslug', 'mysecret', '{clear,refill,withdraw,hold,unhold,refund,transfer,wallet_create,wallet_close,wallet_block,wallet_unblock,wallet_identification,wallet_info,wallet_history,wallet_statistics,transaction_info,wallet_history_by_provider}', '2023-05-17 17:18:55.000000');

insert into public.error_codes (code, description, http_code)
values  ('internal_error', 'ошибка', 400),
        ('no_such_country', 'неправильный код страны', 400),
        ('no_such_currency', 'неправильный код валюты', 400),
        ('validation_err', 'ошибка при парсинге запроса', 400),
        ('hold_amount_notzero', 'заблокированная сумма не 0', 400),
        ('wallet_currency_mismatch', 'указаны кошельки с разными валютами', 400),
        ('wallet_not_active', 'статус не активный', 400),
        ('wallet_locked', 'кошелек занят', 400),
        ('invalid_operation_amount', '', 400),
        ('not_enough_balance', 'в кошельке недостаточно средств', 400),
        ('no_hold_operation', 'по данной транзакции нет операции hold', 400),
        ('operation_already_done', 'по данной транзакции уже произведён unhold', 400),
        ('nothing_to_refund', 'в транзакции нет withdraw или clearing', 400),
        ('wallet_not_blocked', 'кошелек не заблокирован', 400),
        ('transaction_not_found', 'транзакция не найдена', 404),
        ('no_such_identification', 'несуществующий статус идентификации', 400),
        ('phone_already_registered', 'такой телефон уже существует', 400),
        ('order_already_exists', 'для провайдера данного провайдера уже существует данный order_id', 400),
        ('wallet_not_found', 'кошелек не найден', 400),
        ('consumer_not_found', '', 400),
        ('method_not_allowed', '', 400),
        ('consumer_code_taken', 'уже существует consumer с таким кодом', 400);