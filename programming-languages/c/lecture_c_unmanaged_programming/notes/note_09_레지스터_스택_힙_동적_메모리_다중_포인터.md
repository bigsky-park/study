# 레지스터, 스택 & 힙, 동적 메모리, 다중 포인터

이 문서는 Pope Kim [C 언매니지드 프로그래밍](https://www.udemy.com/course/c-unmanaged-programming-by-pocu/) 강의를 듣고 정리한 문서입니다.

## 메모리의 종류

- 프로그램에서 주로 사용하는 부품은 2개이다.
  - CPU
    - 모든 코드의 로직(연산)을 실행하는 제어 장치.
  - 메모리
    - 실행 중인 코드 및 연산의 중간 결과 등을 저장하는 공간.
    - 변수나 배열 등에 대입되는 데이터가 저장된다.
- 메모리는 또 다시 나뉜다. 개발자가 신경써야 하는 수준에서 나누어보자.
  - 스택(stack) 메모리
    - 스택은 특별한 용도(쓰레드)에 사용하라고 별도로 떼어놔 준 것이다.
  - 힙(heap) 메모리
    - 기본은 힙 메모리이다. 힙 메모리는 누구라도 자기 마음대로, 어디에도 사용할 수 있는 메모리이다.
  - 사실 이 둘은 물리적으로는 같은 메모리이다.
- CPU 안에도 저장 공간이 있다.
  - 레지스터: CPU에서만 사용할 수 있는 고속 저장 공간이다.
  - 엄밀히 말하면 메모리는 아니다.

## 스택 메모리

- 함수를 호출할 때 스택에 정확히 어떤 게 어떤 순서로 들어가는지는 함수 호출 규악(calling conveition)에 따라 달라진다. <https://en.wikipedia.org/wiki/Calling_convention>

## 레지스터

- 엄밀한 의미의 메모리는 아니다. 그치만 휘발성으로 데이터를 저장하는 공간이긴 하다.
- 메모리를 읽고 쓰는 게 느린 이유
  - CPU가 메모리에 접근할 때마다 버스를 거쳐야 한다. 즉, CPU가 연산할 때마다 메모리에 접근하는 시간이 발생한다.
  - 대부분의 컴퓨터에 장착하는 메모리는 DRAM(Dynamic Random Access Memory)이다. DRAM은 가격이 저렴하지만 큰 단점이 있다. 기록된 내용을 유지하기 위해서 주기적으로 정보를 다시 써야 한다는 것이다. 빨라지기가 어려운 단점이 있다.
- 그래서...메모리를 읽고 쓰는 건 느리다. 그래서 CPU 전용의 저장 공간을 만들자는 아이디어가 나왔다. 이게 레지스터이다. SRAM을 CPU안에 넣어두는 식이다.
- 레지스터는 CPU가 사용하는 저장 공간 중에 가장 빠른 저장 공간이다.
- CPU가 연산을 할 때 보통 레지스터에 저장되어 있는 데이터를 사용하고 연산 결과도 레지스터에 다시 저장하는 게 보통이다.

## register 키워드

- C에서는 레지스터를 어떻게 사용할까?
  - 레지스터를 사용해달라고 요청은 할 수 있다. 그치만 실제로 그렇게 될지는 컴파일러에 따라 다르다. 따라서 레지스터를 사용해 달라는 '힌트'정도를 주는 의미이다.

  ```c
  int num;
  register size_t i;

  num = 0;

  for (i = 0; i < 1000; ++i)
  {
      num += i;
  }
  printf("num: %d\n", num);

  ```

- register 키워드
  - `register <자료형> <변수명>`
  - 저장 유형 지정자 (storage-class specifier)
  - 가능하다면 해당 변수를 레지스터에 저장할 것을 요청한다. 그러나 실제로 레지스터를 사용할지 말지는 컴파일러가 결정한다.
  - 레지스터는 메모리가 아니다. 따라서 몇 가지 제약 사항이 있다.
    - 변수의 주소를 구할 수 없다.
    - 레지스터 배열을 포인터로 사용 불가.
    - 블록 범위에서만 사용 가능하다. 전역 변수로는 사용할 수 없다.

- 예전 임베디드 시스템에서는 의미가 있었지만, 현재 데스크탑 환경에서는 대부분 의미가 없다. 또한, 요즘은 굳이 register 키워드를 쓰지 않더라도 컴파일러가 잘 최적화를 해준다. 따라서 굳이 프로그래머가 수동으로 사용하지 않는 키워드이다.

## 힙 메모리

- 스택 메모리의 단점
  - 수명
    - 함수가 반환하면 그 안에 있던 데이터가 다 날아간다. 따라서 데이터를 오래 보존하려면 전역 변수 또는 static 키워드를 사용해야 했다.
  - 크기
    - 특정 용도에 쓰라고 별도로 떼어 놓은 메모리이다. 컴파일 시에 결정되는 것으로 너무 크게 잡을 수 없다. 그래서 큰 데이터를 처리해야 할 경우 스택 메모리에 넣지 못 한다.
- 힙 메모리
  - 컴퓨터에 존재하는 범용적 메모리
  - 스택 메모리처럼 특정 용도로 떼어 놓은 것이 아니다.
  - 스택과 달리 컴파일러 및 CPU가 자동으로 메모리 관리를 해주지 않는다. 따라서 프로그래머가 원하는 때에 원하는 만큼 메모리를 할당받아와 사용하고 원할 때 반납할 수 있다.
- 힙 메모리의 장점
  - 용량 제한이 없다. 컴퓨터에 남아있는 메모리만큼 사용이 가능하다.
  - 프로그래머가 데이터의 수명을 직접 제어할 수 있다. 스택에 저장되어 있는 변수처럼 함수 호출이 끝난다고 사라지지 않는다.
- 힙 메모리의 단점
  - 빌려온 메모리를 직접 해제하지 않으면 누구도 그 메모리를 쓸 수 없다.
  - 스택에 비해 할당/해제 속도가 느리다.
    - 스택은 오프셋 개념인데 반해 힙은 사용/비사용 중인 메모리 관리 개념이다.
    - 또한 메모리 공간에 구멍이 생겨 효율적으로 메모리 관리가 어렵기도 하다.
- 스택 메모리는 정적 메모리이다.
  - 이미 공간이 따로 잡혀 있다.
  - 할당/해제가 자동으로 관리되게 코드가 컴파일된다.
  - 오프셋 개념으로 정확히 몇 바이트씩 사용해야 하는지 컴파일 시 결정된다.
- 힙 메모리는 동적 메모리이다.
  - 실행 중에 크기와 할당/해제 시기가 결정된다.

## 동적 메모리

- 동적 메모리를 가져다 사용할 때는 총 세 가지 단계를 거친다.
  - 메모리 할당
    - 프로그램이 힙 관리자에게 메모리를 ~~ 바이트 달라고 요청한다.
  - 메모리 사용
    - 그 메모리를 원하는 대로 사용한다. 이 때까지 포인터 사용했던 것과 다르지 않다.
  - 메모리 해제
    - 프로그램이 힙 관리자에게 메모리를 반환한다.

## 메모리 할당 및 해제 함수, malloc()

- malloc()
  - `void * malloc(size_t size);`
  - 메모리 할당(memory allocation)의 약자
  - size_t 바이트 만큼의 메모리를 반환해준다.
  - 반환된 메모리에 들어있는 값은 쓰레기 값이다. 초기화를 해주지 않는다.
  - 메모리가 더 이상 없는 등 어떤 이유로 실패하면 NULL을 반환한다.

## free(), malloc() 사용 예

- free()
  - 할당받은 메모리를 해제하는 함수
  - 즉, 메모리 할당 함수들을 통해서 얻은 메모리만 해제 가능하다. 그 외의 주소를 매개변수로 전달할 경우 결과가 정의되지 않았다.

- 메모리 할당 예, 간단한 예

```c
#include <stdlib.h>

#define LENGTH (10)

/* ... */

size_t i;
int* nums = malloc(LENGTH * sizeof(int)); /* 메모리 할당 */

for (i = 0; i < LENGTH; ++i)
{
    nums[i] = i * LENGTH;
}

for (i = 0; i < LENGTH; ++i)
{
    printf("%d ", nums[i]);
}

free(nums); /* 중요! 메모리 반환 */

```

- 메모리 할당 예, 여러 줄의 입력을 받아 출력하기

```c
#include <stdlib.h>
#include <stdio.h>
#include <string.h>

#define NUM_LINES (5)
#define LINE_LENGTH (2048)

/* ... */

char* lines[NUM_LINES];
char line[LINE_LENGTH]; /* 일종의 임시 버퍼 */
size_t i;
size_t j;

/* ... */

for (i = 0; i < NUM_LINES; ++i)
{
    if (!fgets(line, LINE_LENGTH, stdin))
    {
        clearerr(stdin);
        break;
    }

    lines[i] = malloc(strlen(line) + 1); /* +1 하는 이유는 널문자를 넣어야 하기 때문 */
    if (lines[i] == NULL)
    {
        fprintf(stderr, "%s\n", "out of memory");
        break;
    }

    strcpy(lines[i], line);
}

for (j = 0; i < i; ++j)
{
    printf("%s, lines[j]);
}

for (j = 0; j < i; ++j)
{
    free(lines[i]); /* 중요 */
}

```

## 동적 메모리 할당 시 문제

- 할당받아 온 주소를 그대로 연산에 사용하면
  - 메모리 할당 함수가 반환한 주소가 저장된 변수를 그대로 포인터 연산에 사용하면 메모리 해제할 때 문제가 발생할 수 있다. 최초에 받아온 주소가 아니라 다른 주소를 가르킬 수 있기 때문이다. 이럴 경우 결과가 정의되지 않고...망할 수 있다.
- 할당받아 온 포인터로 연산을 하지 않는 게 실수를 줄일 수 있는 방법 중 하나이다.

  ```c
  void* nums;
  int* p;
  size_t i;

  nums = malloc(LENGTH * sizeof(int));
  p = nums; /* 다른 변수에 할당 */

  for (i = 0; i < LENGTH; ++i)
  {
      *p++ = 10 * (i + 1);
  }

  free(nums);

  ```

- 해제한 메모리를 다시 해제해도 문제이다. 잘못하면 크래시가 날 수도 있다.
- 해제한 메모리를 또 사용해도 문제가 발생할 수 있다. 결과가 정의되지 않아있긴 하지만 크래시가 날 수 있다.
- 해제한 후 널 포인터를 대입해주는 게 실수를 방지할 수 있는 방법이 될 수 있다.
  - free() 한 뒤에 변수에 NULL을 대입해서 초기화

    ```c
    /* codes... */
    nums = malloc(LENGTH * sizeof(int));

    /* codes... */

    free(nums);
    nums = NULL;

    ```

## free()는 몇 바이트를 해제할지 어떻게 알지?, calloc(), memset(), realloc()

- 구현마다 다르지만 보통 `malloc(32)` 하면 그것보다 조금 큰 메모리를 할당한 뒤, 제일 앞 부분에 어떤 데이터들을 채워 놓는다. 그리고 데이터 만큼의 오프셋을 더한 값을 주소로 돌려준다.
- 나중에 free()를 통해 주소 해제를 요청하면 오프셋만큼 앞으로 다시 가서 그 앞 주소를 본 뒤, 실제 몇 바이트가 할당 됐었는지 확인 후 해제한다. 즉 앞 부분에 일종의 메타데이터를 적어뒀던 것이다.

### 다른 메모리 할당 함수

- calloc()
  - `void* calloc(size_t num, size_t size);`
    - 메모리를 할당할 때 자료형의 크기(size)와 수(num)을 따로 지정한다.
    - 모든 바이트를 0으로 초기화 해준다.
    - 잘 안쓴다. 그 이유는 calloc()을 하나 malloc() + memset()을 하나 똑같은 것이기 때문이다. 보통은 calloc() 대신 malloc()와 memset()을 조합해서 쓴다.

      ```c
      void* nums;

      nums = malloc(LENGTH * sizeof(int));
      memset(nums, 0, LENGTH * sizeof(int));

      free(nums);
      nums = NULL;

      ```

- memset()
  - `void* memset(void* dest, int ch, size_t count);`
  - <string.h>에 있다.
  - char로 초기화(1바이트씩)된다.
  - 그 외의 자료형으로 초기화하려면 직접 for 문을 작성해야 한다.
  - 다음과 같은 경우에 결과가 정의되지 않았다.
    - count가 dest의 영역을 넘어설 경우 (소유하지 않은 메모리에 쓰기)
    - dest가 널 포인터일 경우

- realloc()
  - `void* realloc(void* ptr, size_t new_size);`
  - 이미 존재하는 메모리(ptr)의 크기를 new_size 바이트로 변경한다.
  - 새로운 크기가 허용하는 한 기존 데이터를 그대로 유지한다.

- 크기가 커져야 할 때 두 가지 경우가 있다.
  - 지금 갖고 있는 메모리 뒤에 충분한 공간이 없는 경우, 새로운 메모리를 할당한 뒤 기존 내용을 복사하고 새 주소를 반환한다.

  ```c
  str = malloc(LENGTH * sizeof(char));
  str = realloc(str, 2 * LENGTH * sizeof(char)); /* str은 새로운 주소를 갖는다. */

  ```

  - 지금 갖고 있는 메모리 뒤에 공간이 충분하다면 그냥 기존 주소를 반환한다(보장은 없다). 그리고 추가된 공간을 쓸 수 있게 된다.

  ```c
  str = malloc(LENGTH * sizeof(char));
  str = realloc(str, 2 * LENGTH * sizeof(char)); /* str은 기존 주소를 갖는다. */

  ```  

## realloce()의 메모리 누수 문제, memcpy()

- realloc()을 제대로 안 쓰면 메모리 누수가 발생할 수 있다.
  - `void* realloc(void* ptr, size_t new_size);`
    - 반환값
      - 성공 시, 새롭게 할당된 메모리의 시작 주소를 반환하며 기존 메모리는 해제된다.
      - 실패 시, NULL을 반환하지만 기존 메모리는 해제되지 않는다.
        - 실패 시, 메모리 누수가 발생할 수 있다.

          ```c
          void* nums;

          nums = malloc(SIZE);

          nums = realloc(nums, 2 * SIZE); /* 실패 시, NULL 반환 */

          ```

          - 실패 후 num에 NULL이 대입되어 버린다면, 원래 nums에 저장되어 있던 주소가 사라진다. 따라서 기존 메모리를 해제하고 싶어도 주소를 몰라서 해제할 수 없게 된다. 메모리 누수가 발생하는 것이다.

          ```c
          /* 올바른 방법 */
          void* nums;
          void* tmp;

          nums = malloc(SIZE);

          tmp = realloc(nums, 2 * SIZE);
          if (tmp != NULL)
          {
              nums = tmp;
          }

          ```

- realloc()는 malloc() + memcpy() + free() 를 합친 것과 유사하다.
- memcpy()
  - `void* memcpy(void* dest, const void* src, size_t const);`
    - <string.h>에 있다.
    - src의 데이터를 count 바이트 만큼 dest에 복사한다.
    - 다음과 같은 경우 결과가 정의되지 않는다.
      - dest의 영역 뒤에 데이터를 복사할 경우 (소유하지 않은 메모리에 쓰기)
      - src나 dest가 널 포인터일 경우 (널 포인터 역참조)
- 메모리 누수가 나지 않게 코드를 작성하자
  - realloc()을 사용할 때는 정말 정말 조심해야 한다.
  - 그래서 차라리 malloc() + memcpy() + free()로 좀 더 명시적으로 드러나게 코딩하는 게 나을 수도 있다.
  - 그냥...신경 안 쓰고 realloc()을 쓰는 경우도 많다.

## memcmp(), 정적 vs 동적 메모리

- 여러 줄 입력 받아 출력하기 예

```c
#define LINE_LENGTH (2048)
#define INCREMENT (2)

/* ... */
char** lines;
char line[LINE_LENGTH];
size_t max_lines;
size_t num_lines;
size_t i;
char** tmp;

max_lines = 0;
num_lines = 0;
lines = NULL;

/* ... */

while(1)
{
    if (fgets(line, LINE_LENGTH, stdin) == NULL)
    {
        clearerr(stdin);
        break;
    }

    /* 더 이상 공간이 없음 */
    if (num_lines == max_lines)
    {
        tmp = realloc(lines, (max_lines + INCREMENT) * sizeof(char*));

        if (tmp == NULL)
        {
            fprintf(stderr, "%s\n", "out of memory");
            break;
        }

        lines = tmp;
        max_lines += INCREMENT:
    }

    lines[num_lines] = malloc(strlen(line) + 1);
    if (lines[num_lines] = NULL)
    {
        fprintf(stderr, "%s\n", "out of memory");
        break;
    }
    strcpy(lines[num_lines++], line);
}

for (i = 0; i < num_lines; ++i)
{
    printf("[%d] %s", i, lines[i]);
}

for (i = 0; i < num_lines; ++i)
{
    free(lines[i]);
}

free(lines); /* 포인터를 담을 수 있는 배열도 동적 배열이기 때문에 해제해 주어야 한다. */
return 0;

```

- memcmp()
  - `int memcmp(const void* lhs, const void* rhs, size_t count);`
    - 첫 count 만큼의 메모리를 비교하는 함수
    - strcmp()와 매우 비슷하다.
    - 단, 널 문자를 만나도 계속 진행한다.
    - 다음의 경우 결과가 정의되지 않았다.
      - lhs가 rhs의 크기를 넘어서서 비교할 경우 (소유하지 않은 메모리에 쓰기)
      - lhs이나 rhs이 널 포인터일 경우 (널 포인터 역참조)
  - 구조체를 비교할 때 유용하다.
    - 단 구조체가 포인터 변수를 가질 경우에는 원하는 대로 비교가 동작하지 않을 수 있다.

- 구조체 멤버 변수 - 배열 vs 포인터
  - 고정된 길이인 배열
    - 그대로 대입 가능
    - 파일에 곧바로 저장 가능
    - memcpy()를 곧바로 사용 가능
    - 낭비하는 용량이 있음
    - 메모리 할당/해제 속도 빠름

    ```c
    typedef struct
    {
        char firstname[NAME_LEN];
        char lastname[NAME_LEN];
    } name_fixed_t;

    ```

  - 동적 메모리를 사용하는 포인터
    - 그대로 대입 불가, 이 경우는 얕은 복사가 되어버린다.
    - 파일에 곧바로 저장 불가능
    - memcpy() 곧바로 사용 불가능
    - 낭비하는 용량 없음
    - 메모리 할당/해제 속도가 느리다.

    ```c
    typedef struct
    {
        char* firstname;
        char* lastname;
    } name_dynamic_t;

    ```

- 정적 vs 동적 메모리
  - 정적 메모리를 우선적으로 사용할 것
  - 안 될 때만 동적 메모리를 쓴다.

## 동적 메모리의 소유권 문제

- 이미 할당된 메모리를 누가 해제해 주어야 하는 것인가? 이런 문제들을 정립하는 것이 소유권 문제이다.
- 소유주란, 메모리를 생성한 함수를 말한다. 그리고 그 메모리를 반드시 책임지고 해제해야 하는 주체이다. 소유주가 아닐 경우엔 그냥 빌려서 사용할 뿐 해제하면 안 된다는 말이기도 하다.
- 문제가 생길 수 있는 예 - 호출자가 동적으로 할당한 메모리를 사용하는지 모를 수 있다.

```c
const char* combine_string(const char* a, const char* b)
{
    void* str;
    char* p;

    /* codes */

    str = malloc(size);

    /* codes */

    return str;
}


/* ... */

result = combine_string("Hello", "World");

```

- 이런 문제에 대한 해결 방법은...이런 경우가 최대한 없게 하는 것이다. 함수 안에서 할당하는 대신 함수 밖에서 할당 후 매개변수로 전달하는 것이다.
- 동적으로 할당 후 반환을 피할 수 없다면?
  - 함수 이름, 변수명이나 주석에라도 표시하자
    - 함수 이름: `const char* combile_string_malloc(const char* str1, const char& str2);`
    - 변수 명: `void* pa_str;`, pa: pointer allocated

- 베스트 프랙티스 정리
  - malloc() 작성한 뒤에 곧바로 free()도 추가하자.
  - 동적 할당을 한 메모리 주소를 저장하는 포인터 변수와 포인터 연산에 사용하는 포인터 변수를 분리해 사용하자.
  - 메모리 해제 후 널 포인터를 대입하자.
  - 정적 메모리를 우선으로 사용하고 어쩔 수 없을 때만 동적 메모리를 사용하자.
  - 동적 메모리 할당을 할 경우, 변수와 함수 이름에 그 사실을 명백히 표현하자.

## 다중 포인터, 이중 포인터

- 주소를 저장한 변수의 주소는 어디다가 저장할까? 포인터이다. 그럼 그 포인터에 들어간 데이터의 자료형은?
- 이중 포인터

  ```c
  int num = 10;
  int* p = &num; /* int를 저장하는 포인터 */
  int** pp = &p; /* int 포인터를 저장하는 포인터 */

  ```

- 포인터 변수를 서로 교체하기

```c
void swap(int** n1, int** n2)
{
    int* tmp = n1;

    *n1 = *n2;
    *n2 = tmp;
}

/* ... */

int num1 = 10;
int num2 = 20;

int* p;
int* q;

p = &num1;
q = &num2;

swap(&p, &q);

```
