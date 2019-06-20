int *g;
void func1()
{
    int temp = 0;
    g = &temp;
}

void func2()
{
    func1();
    *g = 3;
}

int main()
{
    func2();
    return 0;
}
