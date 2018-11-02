package quilldelta

import (
	"encoding/json"
	"fmt"
)

type Delta struct {
	Ops   []Op  `json:"ops,omitempty"`
	Attrs Attrs `json:"attributes,omitempty"`
}

func Parse(given string) (Delta, error) {
	var delta Delta
	err := json.Unmarshal([]byte(given), &delta)
	return delta, err
}

const testData = `{"ops":[{"insert":"В британском банке объясняют, что с помощью блокчейн-платформы успешно открыли аккредитив для сельскохозяйственной группы Cargill. По мнению банка, это показывает, что платформа готова для широкого коммерческого внедрения. В качестве преимуществ создатели называют сокращение времени обработки транзакций с дней до часов и отсутствие бумажной волокиты. \\nКо всему прочему, считают в банке, блокчейн-платформы для торговли могут стать единым стандартом. В пример приводят грузовые контейнеры, которые являются универсальной грузовой тарой для железнодорожных, морских и авиаперевозок. Переход на них осуществлялся медленно, но сейчас это стандарт для компаний по всему миру. То же должно произойти и с блокчейн в торговле. \\nТранзакция для Cargill касалась оплаты ставки соевых бобов из Аргентины в Малайзию. HSBC использовала платформу Corda, разработанную блокчейн-консорциумом R3. Голландский банк ING выступил контрагентом по сделке. Как и обещали британцы, сделка прошла без курьеров, бумаги и долгого ожидания. "},{"attributes":{"blockquote":true},"insert":"\\n"},{"insert":"«Следующий этап – это подключение большего числа участников. Например, банки, судоходные компании, порты и таможни должны будут подключиться еще до широкого распространения технологии», – говорит глава отдела инноваций HSBC Вивек Рамачандран.  Он также добавляет, что технология будет активно расти ближайшие пять лет. А успех конкретно этой транзакции открывает дверь на рынок стоимостью $9 трлн.\\nПараллельно HSBC анонсировал и другую технологическую инициативу. В 24 странах банк запустит сервисы распознавания лиц. После этого пользователям не нужно будет запоминать пароли"}]}`

func main() {
	parsed, err := Parse(testData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", parsed)

	bs, _ := json.Marshal(parsed)
	fmt.Println(testData)
	fmt.Println(string(bs))
	fmt.Println(testData == string(bs))
}