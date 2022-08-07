package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = false

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var myText = `В кабинете, полном дыма, шел разговор о войне, которая была объявлена манифестом, о наборе.
Манифеста еще никто не читал, но все знали о его появлении. Граф сидел на оттоманке между двумя курившими и
разговаривавшими соседями. Граф сам не курил и не говорил, а, наклоняя голову то на один бок, то на другой,
с видимым удовольствием смотрел на куривших и слушал разговор двух соседей своих, которых он стравил между собой.
Один из говоривших был штатский, с морщинистым, желчным и бритым худым лицом, человек, уже приближавшийся к старости,
хотя и одетый, как самый модный молодой человек; он сидел с ногами на оттоманке с видом домашнего человека и,
сбоку запустив себе далеко в рот янтарь, порывисто втягивал дым и жмурился. Это был старый холостяк Шиншин,
двоюродный брат графини, злой язык, как про него говорили в московских гостиных. Он, казалось, снисходил до
своего собеседника. Другой, свежий, розовый гвардейский офицер, безупречно вымытый, застегнутый и причесанный,
держал янтарь у середины рта и розовыми губами слегка вытягивал дымок, выпуская его колечками из красивого рта.
Это был тот поручик Берг, офицер Семеновского полка, с которым Борис ехал вместе в полк и которым
Наташа дразнила Веру, старшую графиню, называя Берга ее женихом. Граф сидел между ними и внимательно слушал.
Самое приятное для графа занятие, за исключением игры в бостон, которую он очень любил, было положение слушающего,
особенно когда ему удавалось стравить двух говорливых собеседников.`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})

	t.Run("my test", func(t *testing.T) {
		expected := []string{
			"и",     // 10
			"на",    // 5
			"с",     // 5
			"в",     // 4
			"Граф",  // 3
			"был",   // 3
			"между", // 3
			"не",    // 3
			"о",     // 3
			"он",    // 3
		}
		require.Equal(t, expected, Top10(myText))
	})
}
