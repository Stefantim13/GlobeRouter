#include <bits/stdc++.h>
using namespace std;

struct Zbor
{
    int dest,tdeparture,tarrival,cost
};

vector<vector<Zbor>>ListaZboruri;
vector<pair<int,vector<Zbor>>solutii;
int dest;
int crit;
bool crt()

void BackTrack(int top,int locatie,vector<Zbor>drum)
{
    if(drum[top].dest==dest)
    {
        if(crit==1)
            solutii.push_back(drum[top].tarrival-drum[0].tdeparture,drum);
        else if(crit==2)
        {
            int s=0;
            for(int i=0;i<=top;i++)
                s+=drum[i].cost;
            solutii.push_back(s,drum);
        }
        return;
    }
    if(top>nrEscale)return;
    for(auto x:ListaZboruri[locatie])
    {
        drum[++top]=x;
        BackTrack(top,drum)
        top--;
    }
}


int main()
{

    return 0;
}
