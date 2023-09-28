/operations/refill  
/operations/withdraw  
/operations/confirm  
/operations/hold  
/operations/clear  
/operations/unhold  
/operations/transfer  
/operations/refund       // только для успешных одностадийных или отклиреных средств

/wallet/create  
/wallet/close  
/wallet/block  
/wallet/unblock  
/wallet/info  
/wallet/history  
/wallet/identification  
/wallet/totals  


**4Maintainers.**  
операции можно проводить только на активные кошельки  
В кошельке должно быть достаточно средств на проведение операции

кошелек должен блокироавться на время проведения операции (т.е. не может проходить 2 паралельных звпроса связанные с одним кошельком)

Логику операций оборачиваем в postgres транзакции(RepeatableRead), т.к. проводим несколько апдейтов раздельно и в теории параллельные запросы могли бы портить математику балансов. (atomic_repo)

В теории можно сделать middleware withWallet внутри которого будем сразу искать кошелек.  
Однако придется дважды парсить request.

**TODO**
* Wallet Identification пока существует просто так. 
* Не придумал, как использовать колонку internal_log_message внутри базы.
* /history внутри тупая пагинация, можно сделать через курсор чисто по id т.к. я использую ULID в некоторых таблицах. (см. cmd/migrate/migrations)

# Flow
**controller**
1. Читаем тело из запроса.
2. Блочим учавствующие в запросе кашельки в редисе. (get запросы я не блочу)
3. опракидываем дто в бизнес логику.

**service**
1. валидируем дто.
2. валидируем консумера.
3. проверяем, что все готово к транзакции (активный кошелек, средства и т.д.).
4. кидаем запросы в репозиторий.

**repository** postgres, driver pgx. 
1. Если мы пишем что-то в базу, то обязательно заворачиваем бизнес логику в postgres транзакцию.


# Поднятие. 
1. docker compose up из ./compose 
2. Включаем vault из браузера, кладем туды конфиг по роуту.
3. Подаем полученные ключи в ./config/viper.yaml
4. make migrate
5. make seed 
6. make run