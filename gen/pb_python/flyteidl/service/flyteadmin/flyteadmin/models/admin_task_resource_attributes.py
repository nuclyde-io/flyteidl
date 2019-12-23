# coding: utf-8

"""
    flyteidl/service/admin.proto

    No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)  # noqa: E501

    OpenAPI spec version: version not set
    
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six


class AdminTaskResourceAttributes(object):
    """NOTE: This class is auto generated by the swagger code generator program.

    Do not edit the class manually.
    """

    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'cpu': 'str',
        'gpu': 'str',
        'memory': 'str',
        'storage': 'str'
    }

    attribute_map = {
        'cpu': 'cpu',
        'gpu': 'gpu',
        'memory': 'memory',
        'storage': 'storage'
    }

    def __init__(self, cpu=None, gpu=None, memory=None, storage=None):  # noqa: E501
        """AdminTaskResourceAttributes - a model defined in Swagger"""  # noqa: E501

        self._cpu = None
        self._gpu = None
        self._memory = None
        self._storage = None
        self.discriminator = None

        if cpu is not None:
            self.cpu = cpu
        if gpu is not None:
            self.gpu = gpu
        if memory is not None:
            self.memory = memory
        if storage is not None:
            self.storage = storage

    @property
    def cpu(self):
        """Gets the cpu of this AdminTaskResourceAttributes.  # noqa: E501


        :return: The cpu of this AdminTaskResourceAttributes.  # noqa: E501
        :rtype: str
        """
        return self._cpu

    @cpu.setter
    def cpu(self, cpu):
        """Sets the cpu of this AdminTaskResourceAttributes.


        :param cpu: The cpu of this AdminTaskResourceAttributes.  # noqa: E501
        :type: str
        """

        self._cpu = cpu

    @property
    def gpu(self):
        """Gets the gpu of this AdminTaskResourceAttributes.  # noqa: E501


        :return: The gpu of this AdminTaskResourceAttributes.  # noqa: E501
        :rtype: str
        """
        return self._gpu

    @gpu.setter
    def gpu(self, gpu):
        """Sets the gpu of this AdminTaskResourceAttributes.


        :param gpu: The gpu of this AdminTaskResourceAttributes.  # noqa: E501
        :type: str
        """

        self._gpu = gpu

    @property
    def memory(self):
        """Gets the memory of this AdminTaskResourceAttributes.  # noqa: E501


        :return: The memory of this AdminTaskResourceAttributes.  # noqa: E501
        :rtype: str
        """
        return self._memory

    @memory.setter
    def memory(self, memory):
        """Sets the memory of this AdminTaskResourceAttributes.


        :param memory: The memory of this AdminTaskResourceAttributes.  # noqa: E501
        :type: str
        """

        self._memory = memory

    @property
    def storage(self):
        """Gets the storage of this AdminTaskResourceAttributes.  # noqa: E501


        :return: The storage of this AdminTaskResourceAttributes.  # noqa: E501
        :rtype: str
        """
        return self._storage

    @storage.setter
    def storage(self, storage):
        """Sets the storage of this AdminTaskResourceAttributes.


        :param storage: The storage of this AdminTaskResourceAttributes.  # noqa: E501
        :type: str
        """

        self._storage = storage

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value
        if issubclass(AdminTaskResourceAttributes, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, AdminTaskResourceAttributes):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other
