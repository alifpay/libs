Статусы (Status)

approved - подтверждается, что оплата оплачена
pending - платеж инициирован, ожидает оплаты
failed - Ошибка при проведении оплаты. Не оплачен (отклонен)
ok - запрос проверки, возможность приема оплаты (Check account request)

Список ошибок
Код, Описание, Fatal
104, Транзакция не существует, true
107, Транзакция уже существует(дубликат unknown final status), false
108, Успешно уже проведен(дубликат), true
200, Успешно проведен, true
201, В обработке, false
202, Принято в обработку, false
203, Платеж отклонен, true
301, Check request successful, true
302, Аккаунт найден, true
303, Услуга временно недоступна, операция запрещено, true
304, Больше не пытайтесь, может быть, оплата пройдет много раз, false
305, Недостаточно средств на счету, false
400, Ваш запрос некорректен (ошибка данных, формата), true
404, Аккаунт не найден, true
405, Сумма вне диапазона, true
406, Check request failed, true
500, Внутренняя ошибка сервера, true
503, Сервис недоступен, true
520, Неизвестная ошибка, false
521, Соединение с провайдером не удалось, false
523, Неверный ответ провайдера, false

Признак Fatal означает, что при повторном запросе результат не изменится (ошибка не временная).