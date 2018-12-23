##0、序
题目要求连接k个已排序（Sorted）的链表（Linked List），合并成一条排好序的链表
自然而然得出两种解决方法：

 1. 暴力破解法
 2. 归并排序法
假设每个链表的平均长度是n

两种方法的复杂度都是O（nklogk）

##1、暴力破解法
暴力破解法，顾名思义，就是
1. 直接从k个链表（Linked List）的表头（Head）中按顺序找出最小的元素，抽出来作为新的链表元素
2. 抽出最小元素后，继续遍历这k个链表，直到所有的数都作为最小元素被抽出，新的链表便完成了
看着复杂度似乎挺高，其实可以用优先队列（priority_queue） 
由于每个值都要遍历一遍，一共nk次，每次插入优先队列（priority_queue）需要logk的复杂度，所以时间复杂度是O（nklogk）

代码如下

```
  struct cmp {
      bool operator() (const ListNode* a,const ListNode* b) {
          return a->val  > b->val;
      }
  };
  ListNode* mergeKLists(vector<ListNode*>& lists) {
      priority_queue<ListNode*, vector<ListNode*>, cmp>pq;
      for(auto i:lists) {
          if(i) {
              pq.push(i);
          }
      }
      if(pq.empty()) {
          return nullptr;
      }
      ListNode* ans = pq.top();
      pq.pop();
      ListNode* tail = ans;
      if(tail->next) {
          pq.push(tail->next);
      }
      while(!pq.empty()) {
          tail->next = pq.top();
          tail = tail->next;
          pq.pop();
          if(tail->next) {
              pq.push(tail->next);
          }
      }
      return ans;
  }
```

##2、归并排序法
使用平常的归并排序法，每次把所有链表分成两半，直到递归到只剩两个或一个链表，比较后合并，最后返回就行

分治合并的复杂度是O（klogk）
两条链表合并的复杂度是O（n）
总的时间复杂度是O（nklogk）

代码如下：

```
   ListNode* mergeKLists(vector<ListNode*>& lists) {
       if(lists.size() == 0) {
   		return NULL;
   	}
       return mergeKLists(lists, 0, lists.size() - 1);
   }

   ListNode* mergeKLists(vector<ListNode*>& lists, int left, int right) {
   	
   	if(left == right) {
   		return lists[left];
   	} 
   	int mid = (left + right) / 2;

   	ListNode* Left = mergeKLists(lists, left, mid);
   	ListNode* Right = mergeKLists(lists, mid + 1, right);

   	return mergeTwoLists(Left, Right);
   }

   ListNode* mergeTwoLists(ListNode* Left, ListNode* Right) {
   	ListNode* mergeList = new ListNode(0);
   	if(Left == NULL) {
   		return Right;
   	}
   	if(Right == NULL) {
   		return Left;
   	}

   	if(Left -> val <= Right -> val) {
   		mergeList -> val = Left -> val;
   		mergeList -> next = mergeTwoLists(Left -> next, Right);
   	}
   	else {
   		mergeList -> val = Right -> val;
   		mergeList -> next = mergeTwoLists(Left, Right -> next);
   	}
   	return mergeList;
   }
```
