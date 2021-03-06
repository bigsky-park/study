# 나만의 라이브러리 만들기, C99

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 나만의 라이브러리 만들기

- C에서도 라이브러리를 만들 수 있다.
  - 오브젝트 파일을 모아 라이브러리로 만든다.
  - 다시 컴파일할 필요 없이 코드 재활용이 가능하다.
  - 소스코드 공개 없이(단, 헤더 파일은 예외) 라이브러리 배포가 가능하다.
  - C에서는 두 종류의 라이브러리를 만들 수 있다.
    - 정적 라이브러리
    - 동적 라이브러리

- (복습) 정적 라이브러리와 링크
  - 정적 라이브러리와 링크하는 것을 정적 링킹이라고 한다.
  - 라이브러리 안에 있는 기계어를 최종 실행 파일에 가져다 복사한다.
  - 동적 링킹에 비해,
    - 실행 파일의 크기가 커진다.
    - 메모리를 더 잡아먹을 수 있다.
    - 실행 속도가 빠르다.

## 정적 라이브러리 만들기

- 보통 정적 라이브러리를 사용하는 절차
  - 소스코드들을 컴파일하여 정적 라이브러리를 만든다.
  - 다른 소스 코드들을 작성할 때 위 라이브러리의 헤더 파일들을 사용한다.
  - 컴파일 할 때 정적 라이브러리와 함께 링킹한다.

- (복습) .c 컴파일해서 .o 파일 얻기
  - `clang -std=c89 -W -Wall -pendatic-errors -c simple_math.c -o simple_math.o`
- 정적 라이브러리와 함께 빌드하기
  - `clang -std=c89 -W -Wall -pendatic-errors -I <dir> -L <dir> -l<lib_name> *.c`
    - `-I <dir>`
      - 대문자 i이다. 인클루드 할 때 헤더 파일을 검색할 경로를 추가한다.
    - `-L <dir>`
      - 빌드 시 사용할 라이브러리 파일이 있는 폴더
    - `-l<lib_name>`
      - 소문자 L이다. lib_name 은 빌드 시 사용할 라이브러리이다.
      - 파일명.lib 에서 파일명을 -l 다음에 띄어쓰기 없이 붙인다.

## 동적 라이브러리와 링크

- 동적 라이브러리와 링크하는 것을 동적 링킹이라고 한다.
- 실행 파일 안에 여전히 구멍을 남겨두는 방법이다. 실행 파일을 실행할 때 실제로 링킹이 일어난다.
- 정적 링킹에 비해
  - 실행 파일 크기가 작다.
  - 여러 실행 파일이 동일한 라이브러리를 공유할 수 있다. 메모리가 절약된다.
  - 여러 실행 파일이 이름은 같지만 버전이 다른 동적 라이브러리를 사용한다면...혼란할 수 있다.
- 동적 라이브러리를 사용하는 절차
  - 소스 코드들을 컴파일하여 동적 라이브러리를 만든다.
  - 다른 소스 코드들을 작성할 때 위 라이브러리의 헤더 파일들을 사용한다.
  - 컴파일 할 때 동적 라이브러리와 함께 링킹한다.
    - 단, 동적 라이브러리에 있는 기계어가 실행 파일에 포함되진 않는다.
    - 실행 중에 동적으로 링킹할 수 있는 정보만 포함한다.
    - 따라서 동적 라이브러리 파일도 같이 배포해야 한다.
- 동적 라이브러리와 운영체제
  - 운영체제마다 실행 파일 및 동적 라이브러리 내부 포맷이 다르다.
    - 리눅스 계열: ELF(executable and linkable format) 포맷
    - 윈도우: PE(portable executable) 포맷
  - 운영체제의 동적 링커(dynamic linker)
    - 프로그램이 실행될 때 필요한 동적 라이브러리를 로딩 후 링킹해준다.
    - 동적 라이브러리는 운영체제에 종속된다.
  - 예) 리눅스 계열에서 gcc를 사용할 경우
    - 동적 라이브러리 만들기
      - `gcc -shared <o파일 + 경로> -o <동적 라이브러리 파일명 + 경로>`
      - `gcc <o파일 + 경로> -L<동적 라이브러리 경로> -I<동적 라이브러리 이름> -o <실행파일 이름>`

## 정적 vs 동적 라이브러리

- 동적 라이브러리의 장단점
  - 장점
    - 실행 파일을 바꾸지 않고 동적 라이브러리 파일만 업데이트가 가능하다.
    - 동적 라이브러리 파일을 바꾸지 않고 실행 파일만 업데이트가 가능하다.
    - 필요에 따라 동적 라이브러리를 선택적으로 로딩이 가능하다.
    - 여러 실행 파일들이 같은 동적 라이브러리륵 공유 가능하다.
  - 단점
    - 해킹 당하기 쉽다.
    - DLL 지옥.
- 정적 라이브러리의 장단점
  - 장점
    - 함수의 주소가 공개 안 되니 보다 안전하다.
    - 정확한 버전의 라이브러리가 실행파일 안에 내포되어 있다.
    - 최적화에 유리하다.
  - 단점
    - CPU 세대 별로 실행 파일을 만들어서 배포해야 한다.
    - 라이브러리의 소스코드가 바뀔 때마다 실행파일을 재배포해야 한다.
    - 실행 파일의 크기가 커진다.
    - 실행 중 다른 실행 파일들과 라이브러리 공유가 불가능하다.
- 기본적으로 정적 라이브러리를 사용하자.

## 인라인 함수

- 매크로 함수는 함수 호출의 과부하를 덜어주는 장점이 있으나 디버깅이 어려운 등 단점이 많았다.
- 그런 아쉬움을 해결하기 위해 들어온 것이 인라인 함수이다.
- `inline 반환형 함수_이름(매개변수 목록) {}`
  - 컴파일러에게 최적화해달라고 알려주는 '힌트'
  - 보통 매크로 함수처럼 코드를 복붙해준다. 즉 함수 호출이 사라진다.
  - 그런데 힌트일 뿐이라 컴파일러가 무시할 수도 있다.
  - inline이 없어도 컴파일러가 알아서 최적화를 해줄 수도 있다.
  - 인라인 함수의 구현은 소스 파일이 아니라 헤더 파일에 두어야 한다. 컴파일 할 때 구현을 알 수 있어야 하기 때문이다.

## 인라인 함수 제대로 사용하는 법

- 인라인 함수의 작동법

```c
/* simple_math.h */
inline int add(int op1, int op2)
{
    return op1 + op2;
}

/* main.c */
#include "simple_math.h"

...

int main()
{
    int result = add(10, 20) * add(30, 50);
    printf("%d", result);
    return 0;
}

```

- 매크로 함수와의 차이점
  - 무식하게 코드를 복붙하지 않는다.
  - 인라인 함수는 컴파일러가 컴파일 중에 함수 호출을 코드로 바꿔준다. 결과적으로는 복붙하긴 하지만 좀 더 융통성 있게 잘 복붙할 뿐이다. 컴파일러가 주는 이점을 누릴 수 있다.
- 함수 구현을 알아야 복붙이 가능하다.
  - 트렌스레이션 유닛 안에 인라인 함수의 실제 코드가 있어야 한다.
  - 이 함수의 구현이 다른 c 파일 안에 있으면 불가능하다. c 파일 별로 따로 컴파일되기 때문이다.
  - 따라서 헤더 파일 안에 실제 코드가 있어야 한다.

## 링킹 오류가 나는 이유

- 모든 .o 파일에 같은 함수가 들어 있다. 즉, 동일한 이름의 함수가 여러개 있는 것이다.
- 어떤 함수와 링킹을 해야 하는지 몰라서 오류가 발생한다.
- 이 문제를 해결하기 위해 `inline` 키워드를 사용한다. 컴파일러에게 호출용 함수가 아니라 코드 교체용이라고 알려주는 것이다. 그 결과 링커가 볼 수 있는 함수 심볼을 만들지 않는다.
- 그런데...다른 오류가 날 수 있다.
  - `inline` 그저 힌트일 뿐이다. 컴파일러가 해당 함수를 인라인화 한다는 보장이 없다. 따라서 하지 않으면 문제가 된다.
    - 무슨 문제가 일어날까?
      - 여전히 실행 중에 함수를 호출하겠다는 의미이다.
      - 컴파일 단계에서는 그 함수의 시그니처만 기억한다. 링커가 실제 함수 구현을 찾아 구멍을 메꿔주려고 하는데, 메꿔줄 방법이 없다. 따라서 링킹 오류가 발생한다.
    - 인라인이 안 되면?
      - 함수 심볼이 없다. 따라서 링커 입장에서는 해당 함수를 찾을 방법이 없어 오류를 일으킨다.

## 링킹 오류 해결법

- 해결법 1: 일반 함수도 따로 만든다.
  - 인라인 함수와 똑같이 구현된 일반 함수가 어딘가에 존재하면 된다. 그러나 이런 방식은 코드의 중복을 만든다.
- 해결법 2: `extern`
  - 가장 좋은 방법이다. 코드 중복 없이 함수 하나만 있는 것이다.
  - 인라인이 되면 인라인으로 사용한다.
  - 안 되면 일반 함수로 사용한다.
  - `extern`을 붙이면 링커가 찾을 수 있는 함수도 만들어 준다.
  - 그런데 단순히 extern 만 붙이면 문제가 있다...이 함수를 포함한 헤더를 인클루드 한 .c 파일마다 함수가 생성된다.
- 올바른 해결 법
  - .h 파일 안에 인라인 함수를 만든다.
  - 그에 대응하는 .c 파일을 만든다.
  - 그 파일에서 인라인 함수가 들어 있는 .h 파일을 인클루드 한다.
  - 그 파일에서 인라인 함수를 extern 인라인 함수로 다시 선언한다.

  ```c
  /* simple_math.h */
  #ifndef SIMPLE_MATH_H
  #define SIMPLE_MATH_H

  inline int add(int op1, int op2)
  {
      return op1 + op2;
  }

  #endif /* SIMPLE_MATH_H */

  /* simple_math.c */
  #include "simple_math.h"

  extern inline int add(int op1, int op2);

  ```

## C++ 인라인과의 차이

- C++ 에서는 헤더 파일에 구현한 인라인 함수는 자동적으로 extern이다.
- 따라서 이 헤더 파일을 인클루드 한 .cpp 파일마다 함수 심볼이 생긴다.
- 그러나 표준에 따르면 링커가 이 여러 심볼 중에 하나만 골라서 링킹해야 한다.

- 베스트 프랙티스: 인라인을 쓰자
  - 매크로 함수보다는 인라인이 좋다.
  - 특히 한 줄짜리 코드처럼 매우 간단한 함수일 때 적합하다.
  - 그런데 C에서는 인라인보다는 매크로를 더 자주 사용한다.
  - 일단 인라인이 사용하기 매우 불편하고 헷갈린다.
  - C89 이후 표준에 추가된 기능은 그리 널리 사용되지 않는다.

## restrict 키워드

- 컴파일러가 최적화를 잘 해줄 수 있어 유용한 키워드이다.
- (프로그래머가)포인터 변수 전용인 컴파일러에게 알려주는 힌트이다. 이 포인터 변수의 메모리는 절대 다른 변수와 겹치지 않는다는 의미이다.
  - 메모리 범위가 겹치는 걸 막는 키워드가 아니다.
  - 여전히 범위가 겹치는 포인터를 전달하는 것이 가능하다. 그 경우의 결과는 정의되지 않았다.

```c
int printf(const char* restrict format, ...);
int fprintf(FILE* restrict stream, const char* restrict format, ...);
int sprintf(char* restrict buffer, const char* restrict format, ...);

void* memcpy(void* restrict dest, const void* restrict src, size_t count);

char* strcpy(char* restrict dest, const char* restrict src);

```

## 한 줄 주석

- C89 에서는 /**/ 가 주석을 다는 유일한 방법이었다.
- C99에서는 // 로 주석을 달 수 있다.

## 변수 선언

- C89에서는 모든 변수를 반드시 블록 상단에 선언했어야 한다.
- C99에서는 블록 중간에 변수 선언이 가능해짐.

## va_copy()

- `va_copy(dest, src)`
  - C99에 추가되었다.
  - 가변 인자 목록을 복사하는 매크로 함수이다.
  - dest를 다 사용한 후에는 반드시 va_end()를 호출해야 한다.

## snprintf()

- sprintf()의 문제점
  - `int sprintf(char* buffer, const char* format, ...);`
  - 안전하지 않다. 길이를 받지 않는다. buffer의 크기보다 긴 문자열이 들어와도 중간에 멈추지 않는다. 즉, buffer 범위를 넘어서서 계속 쓴다.
- `int snprintf(char* restrict buffer, size_t bufsz, const char* restrict format, ...);`
  - 최대 bufsz - 1 개의 문자열을 출력한다.
  - 나머지 하나는 널 문자 용이다.

```c
#include <stdio.h>

#define LENGTH (20)

int main()
{
    char buffer[LENGTH];
    const char* name = "Caterina Hassinger";
    int score = 100;

    snprintf(buffer, LENGTH, "%s's score: %d\n", name, score);
    buffer[LENGTH - 1] = '\0';

    return 0;
}

```

- 왜 마지막에 널 문자를 넣어야 할까?
  - C89에서 자기만의 _snprintf() 를 제공한 컴파일러가 있었다. 이 함수는 널 문자를 안 붙여줬다. 그러나 호환 때문에 이 함수가 남아 있는 경우가 있다. 이 함수와 헷갈려 실수하는 것을 방지하기 위해 널 문자를 넣는 코드를 많이 작성한다.

## 새로운 자료형 long long int

- C89 에서는 최소 64비트인 정수형은 없었다. C99에서 생겼다.
- 최소 64비트이고 long 이상의 크기이다. 다른 언어에서는 보통 그냥 long 이다.
- 리터럴은 부호가 있는 경우 `ll, LL` 이고 부호가 없는 경우 `ull, ULL` 이다.
- 서식 문자는 부호가 있는 경우 `%lli` 이고 부호가 없는 경우 `%llu` 이다.

## 불(bool)형

- 두 가지 방법이 있다. `_Bool` 와 `bool (헤더 인클루드 필요)`이다.
- 언어 자체의 지원이라기 보단 헤더 파일을 통한 지원 정도이다.

- `_Bool`
  - 거짓이면 0, 참이면 1이다.
    - char/int/float 와 같은 값을 _Bool 에 넣을 경우, 0에 해당하면 0으로 넣고 그렇지 않으면 전부 1로 변환한다.
- `bool, true, false`
  - `<stdbool.h>` 헤더에 정의 되어 있다.
    - bool: _Bool 을 다시 정의했다.
    - true: 1로 정의
    - false: 0으로 정의

## 고정 폭 정수형

- `<stdint.h>` 헤더에 정의되어 있다.
  - int8_t / uint8_t
  - int16_t / uint16_t
  - int32_t / uint32_t
  - int64_t / uint64_t
- 최솟값과 최댓값
  - 최솟값
    - INT8_MIN
    - INT16_MIN
    - INT32_MIN
    - INT64_MIN
  - 최댓값
    - INT8_MAX / UNIT8_MAX
    - INT16_MAX / UINT16_MAX
    - INT32_MAX / UINT32_MAX
    - INT64_MAX / UINT64_MAX

## 허수를 표현하는 자료형

- _Imaginary
  - 허수를 나타내는 키워드이다.
  - 일부 컴파일러는 지원 안 할 수 있다.
- _Complexy
  - 복소수를 나타내는 키워드이다.
  - 일부 컴파일러는 다른 이름을 사용할 수 있다.
- `<complext.h>`
  - 허수와 복소수와 관련 있는 헤더 파일이다.
  - _Imaginary와 _Complexy를 재정의한 매크로를 제공한다.

## IEEE 754 부동 소수점 지원

- C99의 주요 기능 중 하나라 할 수 있다.
- float는 IEEE 754 32비트 부동 소수점이다.
- double은 IEEE 754 64비트 부동 소수점이다.
- long double은 IEEE 754 확장 정밀도(extended precision) 부동 소수점이다.
