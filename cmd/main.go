package main

import (
	"fmt"
	"golang/internal/deftype"
)

// interface include methods Play and Stop of package deftype file gadget
type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func TryOut(player Player) {
	player.Play("Test track")
	player.Stop()
	// задаём исходный тип TapePlayer, тогда отработает ветвь else
	// если удалить = значение типа для player, тогда отработает ветвь ок,
	// т.к. исходный тип будет TapeRecord for recorder,и ок = true
	//player = deftype.TapePlayer{}
	// используется утверждение типа, для преобразования к конкретному типу
	recorder, ok := player.(deftype.TapeRecorder) // содержит значение TypeRecord
	// ок=true if recorder содержит значение TapeRecorder, для которого можно вызвать Record
	// ок предотвращает панику
	if ok {
		recorder.Record()
		// если убрать ветвь else, то при передачи в параметр функции значения
		// TapePlayer, ситуация паники когда ок = false, как бы игнорируется
		// методы Play и Stop вызываются, Record нет.
	} else {
		fmt.Println("Player was not a TapeRecorder")
	}
}

func main() {
	// Обновляем переменную для хранения любого значения, под	держивающего Player.
	var player Player = deftype.TapePlayer{}
	mixtape := []string{"Jessie's Girls", "Whip it", "9 to 5"}
	playList(player, mixtape)
	player = deftype.TapeRecorder{}
	playList(player, mixtape)

	TryOut(deftype.TapeRecorder{}) // вызовет метод
	TryOut(deftype.TapePlayer{})   // выводит сообщение о несоответствии
	// конкретного типа исходному, с помощью 2-го возвр-го значения ок предотвращается паника

}
