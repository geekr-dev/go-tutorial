package main

import "fmt"

type Student struct {
    id    uint
    name  string
    male  bool
    score float64
}

func NewStudent(id uint, name string, score float64) *Student {
    return &Student{id: id, name: name, score: score}
}

func NewStudentV2(id uint, name string, score float64) Student {
    return Student{id: id, name: name, score: score}
}

func (s Student) GetName() string {
    return s.name
}

func (s *Student) SetName(name string) {
    s.name = name
}

func (s Student) String() string {
    return fmt.Sprintf("{id: %d, name: %s, male: %t, score: %f}",
        s.id, s.name, s.male, s.score)
}

func main() {
    s := NewStudent(1, "学院君", 100)
    s.SetName("学院君1号")   // ok 正常调用指针方法
    fmt.Println(s.GetName()) // ok 调用值方法自动解引用: (*s).GetName()

    s2 := NewStudentV2(2, "学院君", 90)
    s2.SetName("学院君2号")   // ok s2 是可寻址的，所以实际调用: (&s2).SetName("学院君2号")
    fmt.Println(s2.GetName()) // ok 正常调用值方法

    NewStudent(3, "学院君", 80).SetName("学院君3号") // ok 正常调用指针方法
    // NewStudentV2(4, "学院君", 99).SetName("学院君4号") // err 值类型调用指针方法
}
