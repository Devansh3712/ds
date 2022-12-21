#include <iostream>

template <class T>
class SinglyLinkedList {
private:
    struct Node {
        T data;
        Node* next;
        Node(T data) {
            this->data = data;
            next = nullptr;
        };
    };
    Node* root;
public:
    SinglyLinkedList() {
        root = nullptr;
    };
    void insert(T data) {
        Node* newNode = new Node(data);
        if (root == nullptr) {
            root = newNode;
            return;
        };
        Node* temp = root;
        while (temp->next != nullptr)
            temp = temp->next;
        temp->next = newNode;
    };
    void remove(T data) {
        if (root == nullptr)
            return;
        Node* temp = root;
        if (temp->data == data) {
            root = root->next;
            delete temp;
        };
        while (temp != nullptr) {
            if (temp->next->data == data) {
                Node* newTemp = temp->next;
                temp->next = temp->next->next;
                delete newTemp;
                return;
            };
            temp = temp->next;
        };
    };
    bool search(T data) {
        if (root == nullptr)
            return false;
        Node* temp = root;
        while (temp != nullptr) {
            if (temp->data == data)
                return true;
            temp = temp->next;
        };
        return false;
    };
    void reverse() {
        if (root == nullptr || root->next == nullptr)
            return;
        Node* prev = nullptr;
        Node* curr = root;
        Node* next = nullptr;
        while (curr != nullptr) {
            next = curr->next;
            curr->next = prev;
            prev = curr;
            curr = next;
        };
        root = prev;
    };
    void print() {
        if (root == nullptr)
            return;
        Node* temp = root;
        while (temp != nullptr) {
            std::cout << temp->data << " ";
            temp = temp->next;
        }
        std::cout << std::endl;
    };
};

int main() {
    SinglyLinkedList<int> l;
    for (int i = 0; i < 10; i++)
        l.insert(i);
    l.reverse();
    l.print();
    return 0;
};
