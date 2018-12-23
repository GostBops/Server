##0、序
根据两种遍历还原二叉树是一道经典题型
105是根据前序遍历和中序遍历还原二叉树
106是根据后序遍历和中序遍历还原二叉树
原理相同，首先分析一下前、中、后序的遍历方式


##1、前序、中序、后序遍历
先看一下代码实现：
```C++
void order(TreeNode* root) {
    cout << root->val << endl; //preorder
	if (root->left != NULL) postorder(root->left);
	cout << root->val << endl; //inorder
	if (root->right != NULL) postorder(root->right);
	cout << root->val << endl; //postorder
}
```

从代码可以看出三种遍历只有输出顺序的区别

换成二叉树来理解就是：

- 前序遍历先访问了根节点，然后遍历左子树，直到遍历完叶子结点，继续遍历右子树。
- 中序遍历则先遍历左子树，再访问根节点，接着遍历右子树
- 后序遍历则遍历左子树、右子树，最后访问根节点

##2、根据前序遍历和中序遍历还原二叉树
理解了前序遍历和中序遍历，这道题就很容易了。

1. 首先前序遍历数组（preorder_vector）的第一项便是根节点（root），因为它是最先被访问的。
2. 遍历中序遍历数组（inorder_vector），找到根节点（root）的索引，则根节点左边的数组便是根节点的左子树，右边的数组便是根节点的右子树。
3. 根据左子树的长度（size）和右子树的长度（size）即可确定前序遍历数组（preorder_vector）中根节点（root）左子树和右子树的索引范围。
4. 递归重复123步，直到遍历到叶子结点，结束递归。

代码如下：（105题）
```C++
TreeNode * buildTree(vector<int>& preorder, vector<int>& inorder) {
	if (preorder.size() == 0)
		return NULL;
	TreeNode *root = new TreeNode(preorder[0]);
	if (preorder.size() == 1) {
		return root;
	}

	else {
		int root_index;
		for (int i = 0; i < inorder.size(); i++) {
			if (inorder[i] == preorder[0]) {
				root_index = i; //根节点索引
			}
		}
		vector<int> pre_left(&preorder[1], &preorder[1 + root_index]); //左子树前序遍历范围索引
		vector<int> in_left(&inorder[0], &inorder[root_index]); //左子树中序遍历范围索引
		root->left = buildTree(pre_left, in_left);
		int len = preorder.size();
		vector<int> pre_right(preorder.begin() + 1 + root_index, preorder.end()); //右子树前序遍历范围索引
		vector<int> in_right(inorder.begin() + root_index + 1, inorder.end()); //右子树中序遍历范围索引
		root->right = buildTree(pre_right, in_right);
		return root;
	}
}
```


##3、根据后序遍历和中序遍历还原二叉树
与前序遍历和中序遍历的算法基本相同
1. 首先后序遍历数组（postorder_vector）的最后一项便是根节点（root），因为它是最先被访问的。
2. 遍历中序遍历数组（inorder_vector），找到根节点（root）的索引，则根节点左边的数组便是根节点的左子树，右边的数组便是根节点的右子树。
3. 根据左子树的长度（size）和右子树的长度（size）即可确定后序遍历数组（postorder_vector）中根节点（root）左子树和右子树的索引范围。
4. 递归重复123步，直到遍历到叶子结点，结束递归。

代码如下：（106题）
```C++
    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        if (postorder.size() == 0)
			return NULL;
		TreeNode *root = new TreeNode(postorder[postorder.size() - 1]);
		if (postorder.size() == 1) {
			return root;
		}

		else {
			int root_index;
			for (int i = 0; i < inorder.size(); i++) {
				if (inorder[i] == root -> val) {
					root_index = i; //根节点索引
				}
			}
			vector<int> post_left(&postorder[0], &postorder[root_index]); //左子树后序遍历范围索引
			vector<int> in_left(&inorder[0], &inorder[root_index]); //左子树中序遍历范围索引
			root->left = buildTree(in_left, post_left);
			int len = postorder.size();
			vector<int> post_right(postorder.begin() + root_index, postorder.end() - 1); //右子树后序遍历范围索引
			vector<int> in_right(inorder.begin() + root_index + 1, inorder.end());// 右子树中序遍历范围索引
			root->right = buildTree(in_right, post_right);
			return root;
		}
    }
```

##4、根据前序遍历和后序遍历还原二叉树
1. 这个方案是不可行的，因为没有中序遍历数组（inorder_vector），无法根据根节点（root）确定左右子树的索引范围，即无法确定左右子树。
2. 另外考虑这样一种情况，每一个根节点只有左子树或者右子树，不存在同时拥有左右子树的根节点。那么这样构造出来的不同的两棵树的前序遍历（preorder）和后序遍历（postorder）都是相同的，因此无法根据前序遍历（preorder）和后序遍历（postorder）还原唯一的二叉树。

