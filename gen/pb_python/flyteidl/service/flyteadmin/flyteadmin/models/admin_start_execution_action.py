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

from flyteadmin.models.admin_execution_runtime_metadata import AdminExecutionRuntimeMetadata  # noqa: F401,E501
from flyteadmin.models.admin_execution_system_overrides import AdminExecutionSystemOverrides  # noqa: F401,E501
from flyteadmin.models.core_compiled_workflow_closure import CoreCompiledWorkflowClosure  # noqa: F401,E501
from flyteadmin.models.core_literal_map import CoreLiteralMap  # noqa: F401,E501
from flyteadmin.models.core_workflow_execution_identifier import CoreWorkflowExecutionIdentifier  # noqa: F401,E501


class AdminStartExecutionAction(object):
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
        'id': 'CoreWorkflowExecutionIdentifier',
        'closure': 'CoreCompiledWorkflowClosure',
        'inputs': 'CoreLiteralMap',
        'exec_metadata': 'AdminExecutionRuntimeMetadata',
        'exec_sys_overrides': 'AdminExecutionSystemOverrides'
    }

    attribute_map = {
        'id': 'id',
        'closure': 'closure',
        'inputs': 'inputs',
        'exec_metadata': 'exec_metadata',
        'exec_sys_overrides': 'exec_sys_overrides'
    }

    def __init__(self, id=None, closure=None, inputs=None, exec_metadata=None, exec_sys_overrides=None):  # noqa: E501
        """AdminStartExecutionAction - a model defined in Swagger"""  # noqa: E501

        self._id = None
        self._closure = None
        self._inputs = None
        self._exec_metadata = None
        self._exec_sys_overrides = None
        self.discriminator = None

        if id is not None:
            self.id = id
        if closure is not None:
            self.closure = closure
        if inputs is not None:
            self.inputs = inputs
        if exec_metadata is not None:
            self.exec_metadata = exec_metadata
        if exec_sys_overrides is not None:
            self.exec_sys_overrides = exec_sys_overrides

    @property
    def id(self):
        """Gets the id of this AdminStartExecutionAction.  # noqa: E501


        :return: The id of this AdminStartExecutionAction.  # noqa: E501
        :rtype: CoreWorkflowExecutionIdentifier
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this AdminStartExecutionAction.


        :param id: The id of this AdminStartExecutionAction.  # noqa: E501
        :type: CoreWorkflowExecutionIdentifier
        """

        self._id = id

    @property
    def closure(self):
        """Gets the closure of this AdminStartExecutionAction.  # noqa: E501


        :return: The closure of this AdminStartExecutionAction.  # noqa: E501
        :rtype: CoreCompiledWorkflowClosure
        """
        return self._closure

    @closure.setter
    def closure(self, closure):
        """Sets the closure of this AdminStartExecutionAction.


        :param closure: The closure of this AdminStartExecutionAction.  # noqa: E501
        :type: CoreCompiledWorkflowClosure
        """

        self._closure = closure

    @property
    def inputs(self):
        """Gets the inputs of this AdminStartExecutionAction.  # noqa: E501


        :return: The inputs of this AdminStartExecutionAction.  # noqa: E501
        :rtype: CoreLiteralMap
        """
        return self._inputs

    @inputs.setter
    def inputs(self, inputs):
        """Sets the inputs of this AdminStartExecutionAction.


        :param inputs: The inputs of this AdminStartExecutionAction.  # noqa: E501
        :type: CoreLiteralMap
        """

        self._inputs = inputs

    @property
    def exec_metadata(self):
        """Gets the exec_metadata of this AdminStartExecutionAction.  # noqa: E501


        :return: The exec_metadata of this AdminStartExecutionAction.  # noqa: E501
        :rtype: AdminExecutionRuntimeMetadata
        """
        return self._exec_metadata

    @exec_metadata.setter
    def exec_metadata(self, exec_metadata):
        """Sets the exec_metadata of this AdminStartExecutionAction.


        :param exec_metadata: The exec_metadata of this AdminStartExecutionAction.  # noqa: E501
        :type: AdminExecutionRuntimeMetadata
        """

        self._exec_metadata = exec_metadata

    @property
    def exec_sys_overrides(self):
        """Gets the exec_sys_overrides of this AdminStartExecutionAction.  # noqa: E501


        :return: The exec_sys_overrides of this AdminStartExecutionAction.  # noqa: E501
        :rtype: AdminExecutionSystemOverrides
        """
        return self._exec_sys_overrides

    @exec_sys_overrides.setter
    def exec_sys_overrides(self, exec_sys_overrides):
        """Sets the exec_sys_overrides of this AdminStartExecutionAction.


        :param exec_sys_overrides: The exec_sys_overrides of this AdminStartExecutionAction.  # noqa: E501
        :type: AdminExecutionSystemOverrides
        """

        self._exec_sys_overrides = exec_sys_overrides

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
        if issubclass(AdminStartExecutionAction, dict):
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
        if not isinstance(other, AdminStartExecutionAction):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        return not self == other