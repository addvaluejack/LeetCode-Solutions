class Solution{
public:
    string multiply(string num1, string num2) {
        char num1_array[110] = { 0 }, num2_array[110] = { 0 }, result_array[220] = { 0 };

        for (int i = 0; i < num1.length(); i++) {
            num1_array[110 - 1 - i] = num1[num1.length() - 1 - i];
        }

        for (int i = 0; i < num2.length(); i++) {
            num2_array[110 - 1 - i] = num2[num2.length() - 1 - i];
        }

        for (int i = 0; i < num1.length(); i++) {
            for (int j = 0; j < num2.length(); j++) {
                int k = i + j;
                int carry = result_array[220 - 2 - k] > '0' ? int(result_array[220 - 2 - k] - '0') : 0;

                int tmp = (num1_array[110 - 1 - i] - '0')*(num2_array[110 - 1 - j] - '0') + carry;
                carry = tmp / 10;
                result_array[220 - 2 - k] = char(tmp % 10 + '0');

                int m = k + 1;
                while (carry > 0) {
                    tmp = (result_array[220 - 2 - m] > '0' ? int(result_array[220 - 2 - m] - '0') : 0) + carry;
                    carry = tmp / 10;
                    result_array[220 - 2 - m] = char(tmp % 10 + '0');
                    m++;
                }
            }
        }

        int i;
        for (i = 0; i < 220; i++) {
            if (result_array[i] != 0) {
                break;
            }
        }

        string result(result_array + i);
        for (int j = 0; j < result.length(); j++) {
            if (result[j] != '0') {
                return result;
            }
        }
        return "0";
    }
};
